use lazy_static::lazy_static;
use std::collections::HashMap;
use std::default::Default;

use regex::Regex;
use serde::{Deserialize, Serialize};

lazy_static! {
    static ref RESOURCE_MAP: HashMap<String, Regex> = {
        let mut m = HashMap::new();
        m.insert(
            "$ipv4".to_owned(),
            Regex::new(r"^(?P<ipv4>[0-9]+.[0-9]+.[0-9]+.[0-9]+)").unwrap(),
        );
        m.insert(
            "$status_code".to_owned(),
            Regex::new(r"^(?P<status_code>[0-9]{3})").unwrap(),
        );
        m
    };
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub struct Rule {
    pub resource: String,
    pub operator: Operator,
    pub target: String,
    pub action: Option<Action>,
}

#[derive(PartialEq, Clone, Serialize, Deserialize, Debug, Default)]
pub enum Variable {
    Ipv4(String),
    Date(String),
    HTTPVerb(String),
    HTTPStatusCode(usize),
    Path(String),
    #[default]
    Unknown,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
#[serde(untagged)]
pub enum Operator {
    Equal(String),
    Different(String),
    GreaterThan(String),
    LessThan(String),
    GreaterOrEqualThan(String),
    LessOrEqualThan(String),
    Or(String),
    And(String),
    In(String),
    NotIn(String),
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "snake_case")]
pub struct Action {
    pub name: String,
    pub parameters: Vec<String>,
}

#[derive(Debug)]
pub enum RuleError {
    ActionError(String),
    ParsingError(String),
    RuntimeError(String),
}

impl Rule {
    pub fn run(&self, message: String) -> Result<bool, RuleError> {
        let resource = Variable::build(self.resource.clone(), Some(message)).unwrap_or_default();
        let target = Variable::build(self.target.clone(), None).unwrap_or_default();

        match self.operator {
            Operator::Equal(_) => Ok(Operator::equal(resource, target)),
            _ => Err(RuleError::ActionError(
                "That operator is not implemented.".to_owned(),
            )),
        }
    }
}

impl Variable {
    fn build(value: String, message: Option<String>) -> Result<Variable, RuleError> {
        if value.starts_with("$") {
            let msg = message.unwrap();
            Variable::build_resource(value, msg)
        } else {
            Variable::build_target(value)
        }
    }

    fn build_target(value: String) -> Result<Variable, RuleError> {
        let resources: Vec<Result<Variable, RuleError>> = RESOURCE_MAP
            .clone()
            .into_iter()
            .map(|(key, regex)| Variable::make(key, &regex, value.clone()))
            .collect();
        match resources.into_iter().next() {
            Some(resource) => resource,
            None => Err(RuleError::ParsingError(
                "Error on building target".to_owned(),
            )),
        }
    }

    fn make(key: String, regex: &Regex, value: String) -> Result<Variable, RuleError> {
        match key.as_str() {
            "$ipv4" => {
                if regex.is_match(value.as_str()) {
                    let found = regex.find(value.as_str()).unwrap();
                    Ok(Variable::Ipv4(found.as_str().to_owned()))
                } else {
                    Err(RuleError::ParsingError(
                        "Error on build variable".to_owned(),
                    ))
                }
            }
            _ => Err(RuleError::ParsingError(
                "This variable is not implemented.".to_owned(),
            )),
        }
    }

    fn build_resource(value: String, message: String) -> Result<Variable, RuleError> {
        let regex = RESOURCE_MAP.get(&value).unwrap();
        Variable::make(value, regex, message)
    }
}

impl Operator {
    fn check_unknown(resource: Variable, target: Variable) -> bool {
        match (resource, target) {
            (Variable::Unknown, Variable::Unknown) => return false,
            (_, Variable::Unknown) => return false,
            (Variable::Unknown, _) => return false,
            _ => return true,
        }
    }
    fn equal(resource: Variable, target: Variable) -> bool {
        if Operator::check_unknown(resource.clone(), target.clone()) {
            return resource == target;
        }
        return false;
    }
}
