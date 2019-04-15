#!/usr/bin/env python


import os
import boto3
import yaml
import subprocess


REGION = 'eu-west-1'
ENVIRONMENT = os.getenv('ENVIRONMENT', 'staging')


def get_db_instances(region, page_size=20):
    client = boto3.client('rds', region_name=region)
    marker = ""
    pool = []
    while True:
        for instance in pool:
            yield instance

        if marker is None:
            break
        result = client.describe_db_instances(MaxRecords=page_size,
                                              Marker=marker)
        marker = result.get("Marker")
        pool = result.get("DBInstances")


if __name__ == "__main__":
    em_enabled_instances = {'instances': []}

    for instance in get_db_instances(REGION):
        instance_name = instance['DBInstanceIdentifier']
        if (instance.get('EnhancedMonitoringResourceArn', None) and
                instance_name.endswith(ENVIRONMENT)):
            instance = {'instance': instance_name, 'region': REGION}
            em_enabled_instances['instances'].append(instance)

    with open('/config.yml', 'w') as yaml_file:
        yaml.dump(em_enabled_instances, yaml_file, default_flow_style=False)

    rds_exporter = '/rds_exporter \
                --config.file=/config.yml \
                --log.level="debug" \
                --log.trace'

    subprocess.run(rds_exporter, shell=True)
