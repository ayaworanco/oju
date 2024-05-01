#pragma once
#include <iostream>
#include <filesystem>
#include <vector>

namespace entities
{

  typedef struct {
    std::string name;
    std::string key;
    std::string host;
  } Resource;

  class Config
  {
  public:
    std::vector<Resource> resources;

    Config(std::filesystem::path file_path);
  };

};
