#include <iostream>
#include <core/layers/length.hpp>
#include <argparse/argparse.hpp>

int main(int argc, char *argv[]) {
  // TODO: get the message from somewhere
  argparse::ArgumentParser program("length-layer", "0.0.1");
  program.add_argument("-p", "--payload")
    .help("the payload to compute");

  try {
    program.parse_args(argc, argv);

  } catch (const std::exception &error) {
    std::cerr << error.what() << std::endl;
    std::cerr << program;
    return 1;
  }

  auto input = program.get<std::string>("-p");

  layers::Length length_layer(input);
  //auto encoded = length_layer.run();

  // armazen::Armazen armazen;
  // armazen.save_length(encoded);

  //strategies::Context context(std::make_unique<strategies::Process>());
  //context.run("Temperature (43C) exceeds");
  std::cout << input << std::endl;
  return 0;
}
