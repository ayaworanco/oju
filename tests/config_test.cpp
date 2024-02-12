#include <iostream>
#include <gtest/gtest.h>
#include <core/entities/config.hpp>

TEST(ConfigTest, LoadResourcesFromFile)
{
  std::filesystem::path file_path{"fixtures/config.json"};
  Entities::Config config(file_path);

  EXPECT_EQ(config.resources.size(), 2);
}
