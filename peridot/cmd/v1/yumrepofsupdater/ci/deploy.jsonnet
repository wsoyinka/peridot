local bycdeploy = import 'ci/bycdeploy.jsonnet';
local db = import 'ci/db.jsonnet';
local kubernetes = import 'ci/kubernetes.jsonnet';
local temporal = import 'ci/temporal.jsonnet';
local utils = import 'ci/utils.jsonnet';

bycdeploy.new({
  name: 'yumrepofsupdater',
  replicas: if kubernetes.prod() then 4 else 1,
  dbname: 'peridot',
  backend: true,
  migrate: true,
  migrate_command: ['/bin/sh'],
  migrate_args: ['-c', 'exit 0'],
  legacyDb: true,
  command: '/bundle/yumrepofsupdater',
  image: kubernetes.tag('yumrepofsupdater'),
  tag: kubernetes.version,
  dsn: {
    name: 'YUMREPOFSUPDATER_DATABASE_URL',
    value: db.dsn_legacy('peridot', false, 'yumrepofsupdater'),
  },
  requests: if kubernetes.prod() then {
    cpu: '0.5',
    memory: '1G',
  },
  limits: if kubernetes.prod() then {
    cpu: '3',
    memory: '64G',
  },
  service_account_options: {
    annotations: {
      'eks.amazonaws.com/role-arn': 'arn:aws:iam::893168113496:role/peridot_k8s_role',
    }
  },
  ports: [
    {
      name: 'http',
      containerPort: 45102,
      protocol: 'TCP',
    },
  ],
  health: {
    port: 45102,
  },
  env: [
    {
      name: 'YUMREPOFSUPDATER_PRODUCTION',
      value: if kubernetes.dev() then 'false' else 'true',
    },
    if utils.local_image then {
      name: 'YUMREPOFSUPDATER_S3_ENDPOINT',
      value: 'minio.default.svc.cluster.local:9000'
    },
    if utils.local_image then {
      name: 'YUMREPOFSUPDATER_S3_DISABLE_SSL',
      value: 'true'
    },
    if utils.local_image then {
      name: 'YUMREPOFSUPDATER_S3_FORCE_PATH_STYLE',
      value: 'true'
    },
    if kubernetes.prod() then {
      name: 'YUMREPOFSUPDATER_S3_REGION',
      value: 'us-east-2',
    },
    if kubernetes.prod() then {
      name: 'YUMREPOFSUPDATER_S3_BUCKET',
      value: 'resf-peridot-prod',
    },
    temporal.kube_env('YUMREPOFSUPDATER'),
    $.dsn,
  ],
})
