#include <iostream>
#include <core/entities/config.hpp>
#include <fstream>
#include <nlohmann/json.hpp>
#include <filesystem>
#include <vector>
#include <typeinfo>

namespace fs = std::filesystem;
using namespace nlohmann;

namespace entities
{
  Config::Config(fs::path file_path)
  {
    std::ifstream file(file_path);

    json data = json::parse(file);

    for (auto& resource : data["resources"].items())
    {
      Resource r = {
        resource.value()["name"],
        resource.value()["key"],
        resource.value()["host"]
      };
      resources.push_back(r);
    }
  }
};
