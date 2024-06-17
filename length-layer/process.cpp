#include <iostream>
#include <core/strategies/strategy.hpp>

int main() {
  strategies::Context context(std::make_unique<strategies::Process>());
  context.run();
  return 0;
}
