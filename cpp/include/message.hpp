#pragma once

#include <string>
#include <fmt/core.h>
#include "libmsgagent.h"

namespace message {
    enum level_enum {
        INFO = 0,
        WARN = 1,
        ERROR = 2,
    };

    template<typename... Args>
    inline void info(const std::string &message, const Args &... args) {
        sendMsg(level_enum::INFO, message, args...);
    }

    template<typename... Args>
    inline void warn(const std::string &message, const Args &... args) {
        sendMsg(level_enum::WARN, message, args...);
    }

    template<typename... Args>
    inline void error(const std::string &message, const Args &... args) {
        sendMsg(level_enum::ERROR, message, args...);
    }

    template<typename... Args>
    inline void sendMsg(level_enum level, const std::string &msgTemplate, const Args &... args) {

        std::string msg = fmt::format(msgTemplate, args ...);

        GoString msgTemplateGo = {msgTemplate.c_str(), static_cast<ptrdiff_t>(msgTemplate.size())};
        GoString msgGo = {msg.c_str(), static_cast<ptrdiff_t>(msg.size())};
        SendMessage(level, msgTemplateGo, msgGo);
    }
}