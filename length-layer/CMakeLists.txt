set(LENGTH_LAYER_PROCESS process)
set(LENGTH_LAYER_LAMBDA lambda)

add_executable(${LENGTH_LAYER_PROCESS} process.cpp)
add_executable(${LENGTH_LAYER_LAMBDA} lambda.cpp)

target_compile_features(${LENGTH_LAYER_PROCESS} PRIVATE cxx_std_23)
target_link_libraries(${LENGTH_LAYER_PROCESS} PUBLIC ${OJU_LIBRARY_NAME})

target_compile_features(${LENGTH_LAYER_LAMBDA} PRIVATE cxx_std_23)
target_link_libraries(${LENGTH_LAYER_LAMBDA} PUBLIC ${OJU_LIBRARY_NAME})
