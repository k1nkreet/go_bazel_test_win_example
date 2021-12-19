package test_rule_test

import(
        "testing"
        "github.com/bazelbuild/rules_go/go/tools/bazel_testing"
)

func TestMain(m *testing.M) {
        bazel_testing.TestMain(m, bazel_testing.Args{
                Main: `
-- WORKSPACE --
local_repository(
        name = "go_bazel_test_win_example",
        path = "../go_bazel_test_win_example"
)

-- BUILD.bazel --
load("@go_bazel_test_win_example//:def.bzl", "test_rule")

test_rule(
        name = "hello_world",
        srcs = ["hello_world.cc"],
)

-- hello_world.cc --
#include <iostream>

int main() {
        std::cout << "Hello World!" << std::endl;
}
`,
        })
}

func AssertOutput(t *testing.T, output []byte, expected string) {
        if string(output) != expected {
                t.Fatalf("output of bazel process is invalid.\nExpected: %v\n, Actual: %v\n", expected, string(output))
        }
}

func TestTestRule(t *testing.T) {
        out, err := bazel_testing.BazelOutput("run", "//:hello_world")
        if err != nil {
                t.Fatal(err)
        } else {
                AssertOutput(t, out, "Hello World!\n")
        }
}
