add_library(message)

target_sources(message
        PRIVATE
        message.cc
)

find_library(MSGAGENT
        msgagent
)

message(STATUS "MSGAGENT is set to " ${MSGAGENT})

target_include_directories(message
        PUBLIC
        ../include
        ../lib
)

target_link_libraries(message
        PRIVATE
        ${MSGAGENT}
)