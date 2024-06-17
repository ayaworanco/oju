file(GLOB HEADER_LIST
  CONFIGURE_DEPENDS "${oju_SOURCE_DIR}/include/**/*.hpp")
file(GLOB SOURCE_LIST "${oju_SOURCE_DIR}/src/**/*.cpp")

add_library(${OJU_LIBRARY_NAME} ${SOURCE_LIST} ${HEADER_LIST})
target_include_directories(${OJU_LIBRARY_NAME} PUBLIC ../include)
target_compile_features(${OJU_LIBRARY_NAME} PRIVATE cxx_std_23)

FetchContent_Declare(
  nlohmann_json
  GIT_REPOSITORY https://github.com/nlohmann/json.git
  GIT_TAG v3.11.3
)

FetchContent_MakeAvailable(nlohmann_json)

FetchContent_Declare(
  argparse
  GIT_REPOSITORY https://github.com/p-ranav/argparse.git
  GIT_TAG v3.0
)

FetchContent_MakeAvailable(argparse)


target_link_libraries(${OJU_LIBRARY_NAME} nlohmann_json::nlohmann_json argparse::argparse)

source_group(
  TREE "${PROJECT_SOURCE_DIR}/include"
  PREFIX "Header Files"
  FILES ${HEADER_LIST})
