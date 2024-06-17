#include <iostream>
#include <gtest/gtest.h>
#include <core/strategies/strategy.hpp>

TEST(ProcessStrategyTest, Run)
{
  strategies::Context context(std::make_unique<strategies::Process>());
  context.run();
}
