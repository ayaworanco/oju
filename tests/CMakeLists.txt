file(GLOB TEST_LIST "${oju_SOURCE_DIR}/tests/*.cpp")

FetchContent_Declare(
  googletest
  GIT_REPOSITORY https://github.com/google/googletest.git
  GIT_TAG v1.14.0
)

FetchContent_MakeAvailable(googletest)

enable_testing()

# config_test
add_executable(tests ${TEST_LIST})
target_compile_features(tests PRIVATE cxx_std_23)
target_link_libraries(tests PRIVATE core GTest::gtest_main)
add_test(NAME tests COMMAND tests WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR})

include(GoogleTest)
# gtest_discover_tests(config_test)

add_custom_command(
    TARGET tests
    POST_BUILD
    COMMAND ${CMAKE_CTEST_COMMAND} --output-on-failure
    WORKING_DIRECTORY ${CMAKE_BINARY_DIR}
    COMMENT "Running tests..."
)
