#pragma once
#include <iostream>

namespace strategies
{
  class Strategy
  {
    public:
      virtual ~Strategy() = default;
      virtual void run() const = 0;
  };

  class Process : public Strategy
  {
    public:
      void run() const;
  };

  class Lambda : public Strategy
  {
    public:
      void run() const;
  };

  class Context
  {
    /**
     * @var Strategy the context mantains a reference to one strategy object.
     *      The context does not know about the concrete strategy class.
     */
    private:
      std::unique_ptr<Strategy> _strategy;

    public:
      explicit Context(std::unique_ptr<Strategy> &&strategy = {}) : _strategy(std::move(strategy)) {}
      void run() const
      {
        if (_strategy) {
          _strategy->run();
        } else {
          std::cout << "Strategy not defined\n";
        }
      }
  };
}

