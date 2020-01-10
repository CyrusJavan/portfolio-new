const presets = [
    [
        "@babel/react",
        {
            targets: {
                edge: "17",
                firefox: "60",
                chrome: "67",
                safari: "11.1",
            },
            useBuiltIns: "usage",
        },
    ],
    [
        "@babel/env",
        {
            "plugins": [
                ["styled-components", { "ssr": true, "displayName": true, "preprocess": false }]
            ]
        }
    ],
];

module.exports = { presets };