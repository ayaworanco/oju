cmake_minimum_required(VERSION 3.14)

set(OJU_PROJECT_NAME oju)
set(OJU_SOURCE_DIR app)
set(OJU_SOURCE ${OJU_SOURCE_DIR}/main.cpp)
set(CMAKE_CXX_STANDARD 23)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(OJU_PROJECT_VERSION 0.1)
set(OJU_LIBRARY_NAME core)

# Project naming, versioning and description
project(
  ${OJU_PROJECT_NAME}
  VERSION ${OJU_PROJECT_VERSION}
  DESCRIPTION "A tracer/logger using streaming techs"
  LANGUAGES CXX)

if (CMAKE_PROJECT_NAME STREQUAL PROJECT_NAME)
  set(CMAKE_CXX_EXTENSIONS OFF)
  set_property(GLOBAL PROPERTY USE_FOLDERS ON)
  include(CTest)
endif()

include(FetchContent)

# Library code
add_subdirectory(src)

# Executable is here
add_subdirectory(app)


if(CMAKE_PROJECT_NAME STREQUAL PROJECT_NAME AND BUILD_TESTING)
  enable_testing()
  add_subdirectory(tests)
endif()
