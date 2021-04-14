#define NAPI_VERSION 4
#include <napi.h>
#include "build/dql-parse.h"

Napi::String parse(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    if (!info[0].IsString()) {
        Napi::TypeError::New(env, "String expected").ThrowAsJavaScriptException();
    }
    if (!info[1].IsString()) {
        Napi::TypeError::New(env, "String expected").ThrowAsJavaScriptException();
    }

    Napi::String first = info[0].As<Napi::String>();
    Napi::String second = info[1].As<Napi::String>();

    char *result;
    result = (char *) malloc(1);

    Parse(
        const_cast<char*>(first.Utf8Value().c_str()),
        const_cast<char*>(second.Utf8Value().c_str()),
        &result
    );
    return Napi::String::New(env, result);
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports.Set("parse", Napi::Function::New(env, parse));
  return exports;
}

NODE_API_MODULE(NODE_GYP_MODULE_NAME, Init)
