#pragma once
#include <iostream>

namespace layers
{
  class Length
  {
    private:
      std::string log_entry;

    public:
      Length(std::string entry);
      // TODO: a `run` method to return what? (an encoded string for save into a file?)
  };
}
