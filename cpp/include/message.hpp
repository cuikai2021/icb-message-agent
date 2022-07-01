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

        int size_s = std::snprintf(nullptr, 0, msgTemplate.c_str(), args ...) + 1; // Extra space for '\0'
        auto size = static_cast<size_t>( size_s );
        std::unique_ptr<char[]> buf(new char[size]);
        std::snprintf(buf.get(), size, msgTemplate.c_str(), args ...);
        std::string msg = std::string(buf.get(), buf.get() + size - 1); // We don't want the '\0' inside

        GoString msgTemplateGo = {msgTemplate.c_str(), static_cast<ptrdiff_t>(msgTemplate.size())};
        GoString msgGo = {msg.c_str(), static_cast<ptrdiff_t>(msg.size())};
        SendMessage(level, msgTemplateGo, msgGo);
    }
}