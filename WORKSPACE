workspace(
    name = "peridot",
    managed_directories = {"@npm": ["node_modules"]},
)

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# --start python--
load("//wrksp:python_download.bzl", "python_download")

python_download()

load("//wrksp:python_deps.bzl", "python_deps")

python_deps()
# --end python--

http_archive(
    name = "com_google_protobuf",
    sha256 = "3bd7828aa5af4b13b99c191e8b1e884ebfa9ad371b0ce264605d347f135d2568",
    strip_prefix = "protobuf-3.19.4",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.19.4.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.19.4.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "d6b2513456fe2229811da7eb67a444be7785f5323c6708b38d851d2b51e54d83",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.30.0/rules_go-v0.30.0.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
    ],
)

load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_rules_dependencies()

go_register_toolchains(
    nogo = "@peridot//:nogo",
    version = "1.17.7",
)

go_repository(
    name = "org_golang_google_grpc",
    build_file_proto_mode = "disable",
    importpath = "google.golang.org/grpc",
    sum = "h1:weqSxi/TMs1SqFRMHCtBgXRs8k3X39QIDEZ0pRcttUg=",
    version = "v1.44.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:O7DYs+zxREGLKzKoMQrtrEacpb0ZVXA5rIwylE2Xchk=",
    version = "v0.0.0-20220127200216-cd36cc0744dd",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version = "v0.3.0",
)

go_repository(
    name = "com_github_ProtonMail_go_crypto",
    importpath = "github.com/ProtonMail/go-crypto",
    sum = "h1:J2FzIrXN82q5uyUraeJpLIm7U6PffRwje2ORho5yIik=",
    version = "v0.0.0-20220113124808-70ae35bab23f",
    patches = [
      "//patches:0001-Key-ID-subpacket-should-not-be-hashed-or-critical-fo.patch",
    ],
    patch_args = ["-p1"],
)

go_repository(
    name = "org_golang_x_oauth2",
    build_directives = [
        "gazelle:resolve go cloud.google.com/go/compute/metadata @com_google_cloud_go//compute/metadata",
    ],
    importpath = "golang.org/x/oauth2",
    sum = "h1:pkQiBZBvdos9qq4wBAHqlzuZHEXo07pqV06ef90u1WI=",
    version = "v0.0.0-20210514164344-f6687ab2804c",
)

gazelle_dependencies()

load("@io_bazel_rules_go//extras:embed_data_deps.bzl", "go_embed_data_dependencies")

go_embed_data_dependencies()

http_archive(
    name = "build_bazel_rules_nodejs",
    sha256 = "8f5f192ba02319254aaf2cdcca00ec12eaafeb979a80a1e946773c520ae0a2c9",
    urls = ["https://github.com/bazelbuild/rules_nodejs/releases/download/3.7.0/rules_nodejs-3.7.0.tar.gz"],
)

load("@build_bazel_rules_nodejs//:index.bzl", "node_repositories", "yarn_install")

node_repositories(
    node_version = "16.2.0",
    package_json = ["//:package.json"],
    yarn_version = "1.22.10",
)

yarn_install(
    name = "npm",
    package_json = "//:package.json",
    yarn_lock = "//:yarn.lock",
)

# --start docker--
load("//wrksp:docker_download.bzl", "docker_download")

docker_download()

load("//wrksp:docker_deps.bzl", "docker_deps")

docker_deps()
# --end docker--

# --start openapitools--
load("//wrksp:openapitools_download.bzl", "openapitools_download")

openapitools_download()

load("//wrksp:openapitools_deps.bzl", "openapitools_deps")

openapitools_deps()
# --end openapitools--

# --start grpc_gateway--
http_archive(
    name = "com_github_grpc_ecosystem_grpc_gateway_v2",
    sha256 = "c405cd8f1fb54761933c20be2ad02cce4f9a498d7bdb92a5d3789704884b8360",
    strip_prefix = "grpc-gateway-2.7.3",
    urls = [
        "https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/tags/v2.7.3.tar.gz",
    ],
)

load("@com_github_grpc_ecosystem_grpc_gateway_v2//:repositories.bzl", "go_repositories")

go_repositories()
# --end grpc_gateway--

# --start protoc_gen_validate--
http_archive(
    name = "com_envoyproxy_protoc_gen_validate",
    sha256 = "51ba05210a1a2940530455e01c010daa26d504f4b14855a452716772ea39090c",
    strip_prefix = "protoc-gen-validate-0.6.3",
    urls = [
        "https://github.com/envoyproxy/protoc-gen-validate/archive/v0.6.3.tar.gz",
    ],
)

load("@com_envoyproxy_protoc_gen_validate//:dependencies.bzl", "go_third_party")

go_third_party()
# --end protoc_gen_validate--

# --start jsonnet--
http_archive(
    name = "io_bazel_rules_jsonnet",
    sha256 = "d20270872ba8d4c108edecc9581e2bb7f320afab71f8caa2f6394b5202e8a2c3",
    strip_prefix = "rules_jsonnet-0.4.0",
    urls = ["https://github.com/bazelbuild/rules_jsonnet/archive/0.4.0.tar.gz"],
)

load("@io_bazel_rules_jsonnet//jsonnet:jsonnet.bzl", "jsonnet_repositories")

jsonnet_repositories()

load("@google_jsonnet_go//bazel:repositories.bzl", "jsonnet_go_repositories")

jsonnet_go_repositories()

load("@google_jsonnet_go//bazel:deps.bzl", "jsonnet_go_dependencies")

jsonnet_go_dependencies()
# --end jsonnet--

# --start atlassian--
load("//wrksp:atlassian_download.bzl", "atlassian_download")

atlassian_download()

load("//wrksp:atlassian_deps.bzl", "atlassian_deps")

atlassian_deps()
# --end atlassian--

new_local_repository(
    name = "raw_ts_library",
    build_file = "//rules_raw_ts_library:BUILD",
    path = "rules_raw_ts_library",
)

load("//bases/bazel:containers.bzl", "containers")

containers()
