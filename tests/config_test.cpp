#include <iostream>
#include <gtest/gtest.h>
#include <core/entities/config.hpp>

TEST(ConfigTest, LoadResourcesFromFile)
{
  std::filesystem::path file_path{"fixtures/config.json"};
  entities::Config config(file_path);

  EXPECT_EQ(config.resources.size(), 2);
  entities::Resource resource = config.resources[0];
  EXPECT_EQ(resource.name, "test1");
  EXPECT_EQ(resource.host, "test1-host");
  EXPECT_EQ(resource.key, "test1-key");
}
