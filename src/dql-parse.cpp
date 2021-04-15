#define NAPI_VERSION 4
#include <napi.h>
#include "../lib/dql/parse.h"

Napi::String parse(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    return Napi::String::New(env, Parse(
        const_cast<char*>(info[0].As<Napi::String>().Utf8Value().c_str()),
        const_cast<char*>(info[1].As<Napi::String>().Utf8Value().c_str())
    ));
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set("parse", Napi::Function::New(env, parse));
  return exports;
}

NODE_API_MODULE(NODE_GYP_MODULE_NAME, Init)
