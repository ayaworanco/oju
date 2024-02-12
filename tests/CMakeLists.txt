FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG v1.14.0
)

FetchContent_MakeAvailable(googletest)

enable_testing()

# config_test
add_executable(config_test config_test.cpp)
target_compile_features(config_test PRIVATE cxx_std_17)
target_link_libraries(config_test PRIVATE core GTest::gtest_main)
add_test(NAME config_test COMMAND config_test WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR})

include(GoogleTest)
# gtest_discover_tests(config_test)
