#include "message.hpp"

#include <string>

namespace message {

    template<typename... Args>
    void sendMsg(level_enum level, const std::string &msgTemplate, const Args &... args) {
        std::string msg = string_format(msgTemplate, args ...);
        SendMessage(level, buildGoString(msgTemplate.c_str(), msgTemplate.size()),
                       buildGoString(msg.c_str(), msg.size()));
    }

    template<typename... Args>
    std::string string_format(const std::string &format, const Args &... args) {
        int size_s = std::snprintf(nullptr, 0, format.c_str(), args ...) + 1; // Extra space for '\0'
        auto size = static_cast<size_t>( size_s );
        std::unique_ptr<char[]> buf(new char[size]);
        std::snprintf(buf.get(), size, format.c_str(), args ...);
        return std::string(buf.get(), buf.get() + size - 1); // We don't want the '\0' inside
    }

    GoString buildGoString(const char *p, size_t n) {
        return {p, static_cast<ptrdiff_t>(n)};
    }
}

