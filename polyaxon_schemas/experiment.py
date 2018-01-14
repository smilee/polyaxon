# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function

from marshmallow import Schema, fields, post_load, post_dump, validate

from polyaxon_schemas.base import BaseConfig
from polyaxon_schemas.utils import UUID, humanize_timedelta


class ExperimentJobSchema(Schema):
    sequence = fields.Int(allow_none=True)
    uuid = UUID()
    unique_name = fields.Str(allow_none=True)
    role = fields.Str(allow_none=True)
    experiment = UUID()
    experiment_name = fields.Str()
    last_status = fields.Str(allow_none=True)
    is_running = fields.Bool(allow_none=True)
    is_done = fields.Bool(allow_none=True)
    created_at = fields.LocalDateTime()
    updated_at = fields.LocalDateTime()
    started_at = fields.LocalDateTime(allow_none=True)
    finished_at = fields.LocalDateTime(allow_none=True)
    total_run = fields.Str(allow_none=True)
    definition = fields.Dict()

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentJobConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentJobConfig.remove_reduced_attrs(data)


class ExperimentJobConfig(BaseConfig):
    SCHEMA = ExperimentJobSchema
    IDENTIFIER = 'ExperimentJob'
    REDUCED_ATTRIBUTES = [
        'last_status', 'is_running', 'is_done', 'started_at', 'finished_at', 'total_run']
    REDUCED_LIGHT_ATTRIBUTES = ['definition', 'experiment', 'sequence', 'unique_name', 'updated_at']
    DATETIME_ATTRIBUTES = ['created_at', 'updated_at', 'started_at', 'finished_at']

    def __init__(self,
                 uuid,
                 experiment,
                 experiment_name,
                 created_at,
                 updated_at,
                 definition,
                 unique_name=None,
                 sequence=None,
                 role=None,
                 last_status=None,
                 is_running=None,
                 is_done=None,
                 started_at=None,
                 finished_at=None,
                 total_run=None):
        self.uuid = uuid
        self.unique_name = unique_name
        self.sequence = sequence
        self.role = role
        self.experiment = experiment
        self.experiment_name = experiment_name
        self.created_at = self.localize_date(created_at)
        self.updated_at = self.localize_date(updated_at)
        self.started_at = self.localize_date(started_at)
        self.finished_at = self.localize_date(finished_at)
        self.definition = definition
        self.last_status = last_status
        self.is_running = is_running
        self.is_done = is_done
        self.total_run = None
        if all([self.started_at, self.finished_at]):
            self.total_run = humanize_timedelta((self.finished_at - self.started_at).seconds)


class ExperimentSchema(Schema):
    sequence = fields.Int(allow_none=True)
    uuid = UUID(allow_none=True)
    unique_name = fields.Str(allow_none=True)
    user = fields.Str(validate=validate.Regexp(regex=r'^[-a-zA-Z0-9_]+\Z'), allow_none=True)
    project = UUID(allow_none=True)
    project_name = fields.Str(allow_none=True)
    experiment_group = UUID(allow_none=True)
    experiment_group_name = fields.Str(allow_none=True)
    description = fields.Str(allow_none=True)
    last_status = fields.Str(allow_none=True)
    last_metric = fields.Dict(allow_none=True)
    created_at = fields.LocalDateTime(allow_none=True)
    updated_at = fields.LocalDateTime(allow_none=True)
    started_at = fields.LocalDateTime(allow_none=True)
    finished_at = fields.LocalDateTime(allow_none=True)
    total_run = fields.Str(allow_none=True)
    is_running = fields.Bool(allow_none=True)
    is_done = fields.Bool(allow_none=True)
    is_clone = fields.Bool(allow_none=True)
    config = fields.Dict(allow_none=True)
    content = fields.Str(allow_none=True)
    num_jobs = fields.Int(allow_none=True)
    jobs = fields.Nested(ExperimentJobSchema, many=True, allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentConfig.remove_reduced_attrs(data)


class ExperimentConfig(BaseConfig):
    SCHEMA = ExperimentSchema
    IDENTIFIER = 'Experiment'
    REDUCED_ATTRIBUTES = [
        'user', 'sequence', 'description', 'config', 'jobs', 'content',
        'created_at', 'updated_at', 'started_at', 'finished_at',
        'is_clone', 'is_running', 'is_done', 'total_run', 'last_metric']
    REDUCED_LIGHT_ATTRIBUTES = [
        'uuid', 'project', 'experiment_group', 'description', 'config', 'content',
        'jobs', 'updated_at'
    ]
    DATETIME_ATTRIBUTES = ['created_at', 'updated_at', 'started_at', 'finished_at']

    def __init__(self,
                 sequence=None,
                 user=None,
                 uuid=None,
                 unique_name=None,
                 project=None,
                 project_name=None,
                 experiment_group=None,
                 experiment_group_name=None,
                 description=None,
                 last_status=None,
                 last_metric=None,
                 created_at=None,
                 updated_at=None,
                 started_at=None,
                 finished_at=None,
                 is_clone=None,
                 is_running=None,
                 is_done=None,
                 config=None,
                 content=None,
                 num_jobs=0,
                 jobs=None,
                 total_run=None):
        self.sequence = sequence
        self.user = user
        self.uuid = uuid
        self.unique_name = unique_name
        self.project = project
        self.project_name = project_name
        self.experiment_group = experiment_group
        self.experiment_group_name = experiment_group_name
        self.description = description
        self.last_status = last_status
        self.last_metric = last_metric
        self.started_at = self.localize_date(started_at)
        self.finished_at = self.localize_date(finished_at)
        self.created_at = self.localize_date(created_at)
        self.updated_at = self.localize_date(updated_at)
        self.is_running = is_running
        self.is_done = is_done
        self.is_clone = is_clone
        self.config = config  # The json compiled content of this experiment
        self.content = content  # The yaml content when the experiment is independent
        self.num_jobs = num_jobs
        self.jobs = jobs
        self.total_run = None
        if all([self.started_at, self.finished_at]):
            self.total_run = humanize_timedelta((self.finished_at - self.started_at).seconds)


class ExperimentStatusSchema(Schema):
    uuid = UUID()
    experiment = UUID()
    created_at = fields.LocalDateTime()
    status = fields.Str()

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentStatusConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentStatusConfig.remove_reduced_attrs(data)


class ExperimentStatusConfig(BaseConfig):
    SCHEMA = ExperimentStatusSchema
    IDENTIFIER = 'ExperimentStatus'
    DATETIME_ATTRIBUTES = ['created_at']
    REDUCED_LIGHT_ATTRIBUTES = ['experiment', 'uuid']

    def __init__(self, uuid, experiment, created_at, status):
        self.uuid = uuid
        self.experiment = experiment
        self.created_at = self.localize_date(created_at)
        self.status = status


class ExperimentMetricSchema(Schema):
    uuid = UUID()
    experiment = UUID()
    created_at = fields.LocalDateTime()
    values = fields.Dict()

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentMetricConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentMetricConfig.remove_reduced_attrs(data)


class ExperimentMetricConfig(BaseConfig):
    SCHEMA = ExperimentMetricSchema
    IDENTIFIER = 'ExperimentMetric'
    DATETIME_ATTRIBUTES = ['created_at']
    REDUCED_LIGHT_ATTRIBUTES = ['experiment', 'uuid']

    def __init__(self, uuid, experiment, created_at, values):
        self.uuid = uuid
        self.experiment = experiment
        self.created_at = self.localize_date(created_at)
        self.values = values


class ExperimentJobStatusSchema(Schema):
    uuid = UUID()
    job = UUID()
    created_at = fields.LocalDateTime()
    status = fields.Str()
    message = fields.Str(allow_none=True)
    details = fields.Dict(allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ExperimentJobStatusConfig(**data)

    @post_dump
    def unmake(self, data):
        return ExperimentJobStatusConfig.remove_reduced_attrs(data)


class ExperimentJobStatusConfig(BaseConfig):
    SCHEMA = ExperimentJobStatusSchema
    IDENTIFIER = 'ExperimentJobStatus'
    REDUCED_ATTRIBUTES = ['message', 'details']
    REDUCED_LIGHT_ATTRIBUTES = ['job', 'details', 'uuid']
    DATETIME_ATTRIBUTES = ['created_at']

    def __init__(self, uuid, job, created_at, status, message=None, details=None):
        self.uuid = uuid
        self.job = job
        self.created_at = self.localize_date(created_at)
        self.status = status
        self.message = message
        self.details = details


class JobLabelSchema(Schema):
    project_name = fields.Str()
    experiment_group_name = fields.Str(allow_none=True)
    experiment_name = fields.Str()
    project_uuid = UUID()
    experiment_group_uuid = UUID(allow_none=True)
    experiment_uuid = UUID()
    task_type = fields.Str()
    task_idx = fields.Str()
    task = fields.Str()
    job_uuid = UUID()
    role = fields.Str()
    type = fields.Str()

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return JobLabelConfig(**data)

    @post_dump
    def unmake(self, data):
        return JobLabelConfig.remove_reduced_attrs(data)


class JobLabelConfig(BaseConfig):
    SCHEMA = JobLabelSchema
    IDENTIFIER = 'JobLabel'
    REDUCED_ATTRIBUTES = ['experiment_group_name', 'experiment_group_uuid']

    def __init__(self,
                 project_name,
                 experiment_name,
                 project_uuid,
                 experiment_uuid,
                 task_type,
                 task_idx,
                 job_uuid,
                 role,
                 type,
                 experiment_group_name=None,
                 experiment_group_uuid=None):
        self.project_name = project_name
        self.experiment_group_name = experiment_group_name
        self.experiment_name = experiment_name
        self.project_uuid = project_uuid
        self.experiment_group_uuid = experiment_group_uuid
        self.experiment_uuid = experiment_uuid
        self.task_type = task_type
        self.task_idx = task_idx
        self.job_uuid = job_uuid
        self.role = role
        self.type = type


class PodStateSchema(Schema):
    event_type = fields.Str()
    labels = fields.Nested(JobLabelSchema)
    phase = fields.Str()
    deletion_timestamp = fields.LocalDateTime(allow_none=True)
    pod_conditions = fields.Dict(allow_none=True)
    container_statuses = fields.Dict(allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return PodStateConfig(**data)

    @post_dump
    def unmake(self, data):
        return PodStateConfig.remove_reduced_attrs(data)


class PodStateConfig(BaseConfig):
    SCHEMA = PodStateSchema
    IDENTIFIER = 'PodState'
    DATETIME_ATTRIBUTES = ['deletion_timestamp']

    def __init__(self,
                 event_type,
                 labels,
                 phase,
                 deletion_timestamp=None,
                 pod_conditions=None,
                 container_statuses=None):
        self.event_type = event_type
        self.labels = labels
        self.phase = phase
        self.deletion_timestamp = self.localize_date(deletion_timestamp)
        self.pod_conditions = pod_conditions
        self.container_statuses = container_statuses


class JobStateSchema(Schema):
    status = fields.Str()
    message = fields.Str(allow_none=True)
    details = fields.Nested(PodStateSchema, allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return JobStateConfig(**data)

    @post_dump
    def unmake(self, data):
        return JobStateConfig.remove_reduced_attrs(data)


class JobStateConfig(BaseConfig):
    SCHEMA = JobStateSchema
    IDENTIFIER = 'JobState'

    def __init__(self, status, message=None, details=None):
        self.status = status
        self.message = message
        self.details = details


class ContainerGPUResourcesSchema(Schema):
    index = fields.Int()
    uuid = fields.Str()
    name = fields.Str()
    minor = fields.Int()
    bus_id = fields.Str()
    serial = fields.Str()
    temperature_gpu = fields.Int()
    utilization_gpu = fields.Int()
    power_draw = fields.Int()
    power_limit = fields.Int()
    memory_free = fields.Int()
    memory_used = fields.Int()
    memory_total = fields.Int()
    memory_utilization = fields.Int()
    processes = fields.List(fields.Dict(), allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ContainerGPUResourcesConfig(**data)

    @post_dump
    def unmake(self, data):
        return ContainerGPUResourcesConfig.remove_reduced_attrs(data)


class ContainerGPUResourcesConfig(BaseConfig):
    SCHEMA = ContainerGPUResourcesSchema
    IDENTIFIER = 'ContainerGPUResources'
    MEM_SIZE_ATTRIBUTES = ['memory_free', 'memory_used', 'memory_total']

    def __init__(self,
                 index,
                 uuid,
                 name,
                 minor,
                 bus_id,
                 serial,
                 temperature_gpu,
                 utilization_gpu,
                 power_draw,
                 power_limit,
                 memory_free,
                 memory_used,
                 memory_total,
                 memory_utilization,
                 processes=None):
        self.index = index
        self.uuid = uuid
        self.name = name
        self.minor = minor
        self.bus_id = bus_id
        self.serial = serial
        self.temperature_gpu = temperature_gpu
        self.utilization_gpu = utilization_gpu
        self.power_draw = power_draw
        self.power_limit = power_limit
        self.memory_free = memory_free
        self.memory_used = memory_used
        self.memory_total = memory_total
        self.memory_utilization = memory_utilization
        self.processes = processes


class ContainerResourcesSchema(Schema):
    job_uuid = UUID()
    experiment_uuid = UUID()
    container_id = fields.Str()
    cpu_percentage = fields.Float()
    percpu_percentage = fields.List(fields.Float(), allow_none=True)
    memory_used = fields.Int()
    memory_limit = fields.Int()
    gpu_resources = fields.Nested(ContainerGPUResourcesSchema, many=True, allow_none=True)

    class Meta:
        ordered = True

    @post_load
    def make(self, data):
        return ContainerResourcesConfig(**data)

    @post_dump
    def unmake(self, data):
        return ContainerResourcesConfig.remove_reduced_attrs(data)


class ContainerResourcesConfig(BaseConfig):
    SCHEMA = ContainerResourcesSchema
    IDENTIFIER = 'ContainerResources'
    PERCENT_ATTRIBUTES = ['cpu_percentage']
    MEM_SIZE_ATTRIBUTES = ['memory_used', 'memory_limit']

    def __init__(self,
                 job_uuid,
                 experiment_uuid,
                 container_id,
                 cpu_percentage,
                 percpu_percentage,
                 memory_used,
                 memory_limit,
                 gpu_resources=None):
        self.job_uuid = job_uuid
        self.experiment_uuid = experiment_uuid
        self.container_id = container_id
        self.cpu_percentage = cpu_percentage
        self.percpu_percentage = percpu_percentage
        self.memory_used = memory_used
        self.memory_limit = memory_limit
        self.gpu_resources = gpu_resources
