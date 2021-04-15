{
    "targets": [{
        "target_name": "dql_parse",
        "cflags!": [ "-fno-exceptions" ],
        "cflags_cc!": [ "-fno-exceptions" ],
        "sources": [
            "src/dql-parse.cpp"
        ],
        'include_dirs': [
            "<!@(node -p \"require('node-addon-api').include\")"
        ],
        "libraries": [
             "<(module_root_dir)/lib/dql/parse.so"
        ],
        'dependencies': [
            "<!(node -p \"require('node-addon-api').gyp\")"
        ],
        'defines': [ 'NAPI_DISABLE_CPP_EXCEPTIONS' ],
        'conditions': [
            ['OS=="mac"', {
                'cflags+': ['-fvisibility=hidden'],
                'xcode_settings': {
                  'GCC_SYMBOLS_PRIVATE_EXTERN': 'YES', # -fvisibility=hidden
                }
            }]
        ]
    }]
}
