file(GLOB HEADER_LIST
  CONFIGURE_DEPENDS "${oju_SOURCE_DIR}/include/core/*.hpp")

add_library(${OJU_LIBRARY_NAME} entities/config.cpp ${HEADER_LIST})
target_include_directories(${OJU_LIBRARY_NAME} PUBLIC ../include)
target_compile_features(${OJU_LIBRARY_NAME} PRIVATE cxx_std_17)

FetchContent_Declare(
  nlohmann_json
  GIT_REPOSITORY https://github.com/nlohmann/json.git
  GIT_TAG v3.11.3
)

FetchContent_MakeAvailable(nlohmann_json)


target_link_libraries(${OJU_LIBRARY_NAME} nlohmann_json::nlohmann_json)

source_group(
  TREE "${PROJECT_SOURCE_DIR}/include"
  PREFIX "Header Files"
  FILES ${HEADER_LIST})
