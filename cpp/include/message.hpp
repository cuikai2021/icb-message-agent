#pragma once

#include <string>
#include "libmsgagent.h"

namespace message {
    enum level_enum {
        INFO = 0,
        WARN = 1,
        ERROR = 2,
    };

    template<typename... Args>
    void info(const std::string &message, const Args &... args) {
        sendMsg(level_enum::INFO, message, args...);
    }

    template<typename... Args>
    void warn(const std::string &message, const Args &... args) {
        sendMsg(level_enum::WARN, message, args...);
    }

    template<typename... Args>
    void error(const std::string &message, const Args &... args) {
        sendMsg(level_enum::ERROR, message, args...);
    }

    template<typename... Args>
    void sendMsg(level_enum level, const std::string &message, const Args &... args);

    template<typename ... Args>
    std::string string_format(const std::string &format, const Args &... args);

    GoString buildGoString(const char *p, size_t n);
}