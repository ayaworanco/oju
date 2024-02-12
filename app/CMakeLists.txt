add_executable(${OJU_PROJECT_NAME} main.cpp)

target_compile_features(${OJU_PROJECT_NAME} PRIVATE cxx_std_23)
target_link_libraries(${OJU_PROJECT_NAME} PUBLIC ${OJU_LIBRARY_NAME})
