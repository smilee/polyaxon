# Generated by Django 2.2 on 2019-04-18 14:00

import django.contrib.postgres.fields.jsonb
from django.db import migrations, models
from django.db.models import ExpressionWrapper, F
import django.core.validators
import libs.blacklist
import re

import libs.spec_validation


def create_cluster_owner(apps, schema_editor):
    Cluster = apps.get_model('db', 'Cluster')
    Owner = apps.get_model('db', 'Owner')
    ContentType = apps.get_model('contenttypes', 'ContentType')
    cluster_type_id = ContentType.objects.get_for_model(Cluster).id
    for cluster in Cluster.objects.all():
        Owner.objects.create(object_id=cluster.id,
                             content_type_id=cluster_type_id,
                             name=cluster.uuid)


def migrate_experimentgroup_config(apps, schema_editor):
    ExperimentGroup = apps.get_model('db', 'ExperimentGroup')

    ExperimentGroup.objects.filter(backend__isnull=True).update(backend='native')


def migrate_build_jobs_config(apps, schema_editor):
    BuildJob = apps.get_model('db', 'BuildJob')

    BuildJob.objects.update(content=ExpressionWrapper(F('config'), output_field=str))


def migrate_experiments_config(apps, schema_editor):
    Experiment = apps.get_model('db', 'Experiment')

    Experiment.objects.update(content=ExpressionWrapper(F('config'), output_field=str))


def migrate_jobs_config(apps, schema_editor):
    Job = apps.get_model('db', 'Job')

    Job.objects.update(content=ExpressionWrapper(F('config'), output_field=str))


def migrate_notebook_jobs_config(apps, schema_editor):
    NotebookJob = apps.get_model('db', 'NotebookJob')

    NotebookJob.objects.update(content=ExpressionWrapper(F('config'), output_field=str))


def migrate_tensorboard_jobs_config(apps, schema_editor):
    TensorboardJob = apps.get_model('db', 'TensorboardJob')

    TensorboardJob.objects.update(content=ExpressionWrapper(F('config'), output_field=str))


class Migration(migrations.Migration):
    dependencies = [
        ('db', '0020_auto_20190307_1611'),
    ]

    operations = [
        migrations.RenameField(
            model_name='buildjob',
            old_name='in_cluster',
            new_name='is_managed',
        ),
        migrations.RenameField(
            model_name='experiment',
            old_name='in_cluster',
            new_name='is_managed',
        ),
        migrations.RenameField(
            model_name='job',
            old_name='in_cluster',
            new_name='is_managed',
        ),
        migrations.RenameField(
            model_name='notebookjob',
            old_name='in_cluster',
            new_name='is_managed',
        ),
        migrations.RenameField(
            model_name='tensorboardjob',
            old_name='in_cluster',
            new_name='is_managed',
        ),
        migrations.AddField(
            model_name='job',
            name='backend',
            field=models.CharField(blank=True,
                                   help_text='The default backend use for running this entity.',
                                   max_length=16, null=True),
        ),
        migrations.AlterField(
            model_name='buildjob',
            name='backend',
            field=models.CharField(blank=True,
                                   help_text='The default backend use for running this entity.',
                                   max_length=16, null=True),
        ),
        migrations.AlterField(
            model_name='experiment',
            name='backend',
            field=models.CharField(blank=True,
                                   help_text='The default backend use for running this entity.',
                                   max_length=16, null=True),
        ),
        migrations.AddField(
            model_name='experimentgroup',
            name='backend',
            field=models.CharField(blank=True,
                                   help_text='The default backend use for running this entity.',
                                   max_length=16, null=True),
        ),
        migrations.AddField(
            model_name='pipeline',
            name='backend',
            field=models.CharField(blank=True,
                                   help_text='The default backend use for running this entity.',
                                   max_length=16, null=True),
        ),
        migrations.AlterField(
            model_name='buildjob',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AlterField(
            model_name='experiment',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AlterField(
            model_name='job',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AlterField(
            model_name='notebookjob',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AlterField(
            model_name='tensorboardjob',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AddField(
            model_name='experimentgroup',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AddField(
            model_name='pipeline',
            name='is_managed',
            field=models.BooleanField(default=True,
                                      help_text='If this entity is managed by the platform.'),
        ),
        migrations.AlterField(
            model_name='buildjob',
            name='config',
            field=django.contrib.postgres.fields.jsonb.JSONField(
                blank=True,
                help_text='The compiled polyaxonfile for the build job.',
                null=True, validators=[
                    libs.spec_validation.validate_build_spec_config]),
        ),
        migrations.AlterField(
            model_name='job',
            name='config',
            field=django.contrib.postgres.fields.jsonb.JSONField(
                blank=True,
                help_text='The compiled polyaxonfile for the run job.',
                null=True, validators=[
                    libs.spec_validation.validate_job_spec_config]),
        ),
        migrations.AddField(
            model_name='buildjob',
            name='content',
            field=models.TextField(blank=True,
                                   help_text='The yaml content of the polyaxonfile/specification.',
                                   null=True,
                                   validators=[libs.spec_validation.validate_build_spec_config]),
        ),
        migrations.AddField(
            model_name='experiment',
            name='content',
            field=models.TextField(blank=True,
                                   help_text='The yaml content of the polyaxonfile/specification.',
                                   null=True, validators=[
                    libs.spec_validation.validate_experiment_spec_config]),
        ),
        migrations.AddField(
            model_name='job',
            name='content',
            field=models.TextField(blank=True,
                                   help_text='The yaml content of the polyaxonfile/specification.',
                                   null=True,
                                   validators=[libs.spec_validation.validate_job_spec_config]),
        ),
        migrations.AddField(
            model_name='notebookjob',
            name='content',
            field=models.TextField(blank=True,
                                   help_text='The yaml content of the polyaxonfile/specification.',
                                   null=True,
                                   validators=[libs.spec_validation.validate_notebook_spec_config]),
        ),
        migrations.AddField(
            model_name='tensorboardjob',
            name='content',
            field=models.TextField(blank=True,
                                   help_text='The yaml content of the polyaxonfile/specification.',
                                   null=True, validators=[
                    libs.spec_validation.validate_tensorboard_spec_config]),
        ),
        migrations.AlterField(
            model_name='pipelinerunstatus',
            name='status',
            field=models.CharField(blank=True,
                                   choices=[('created', 'created'), ('warning', 'warning'),
                                            ('scheduled', 'scheduled'), ('running', 'running'),
                                            ('done', 'done'), ('failed', 'failed'),
                                            ('upstream_failed', 'upstream_failed'),
                                            ('stopped', 'stopped'), ('succeeded', 'succeeded'),
                                            ('stopping', 'stopping'), ('skipped', 'skipped'),
                                            ('unknown', 'unknown')], default='created',
                                   max_length=64, null=True),
        ),
        migrations.AlterField(
            model_name='buildjob',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='experiment',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='experimentchartview',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='experimentgroup',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='experimentgroupchartview',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='job',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='notebookjob',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='operation',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='pipeline',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='search',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AlterField(
            model_name='tensorboardjob',
            name='name',
            field=models.CharField(blank=True, default=None, max_length=128, null=True, validators=[
                django.core.validators.RegexValidator(re.compile('^[-a-zA-Z0-9_]+\\Z'),
                                                      "Enter a valid 'slug' consisting of letters, numbers, underscores or hyphens.",
                                                      'invalid'),
                libs.blacklist.validate_blacklist_name]),
        ),
        migrations.AddField(
            model_name='buildjob',
            name='valid',
            field=models.NullBooleanField(default=True),
        ),
        migrations.RunPython(migrate_build_jobs_config),
        migrations.RunPython(migrate_experiments_config),
        migrations.RunPython(migrate_jobs_config),
        migrations.RunPython(migrate_notebook_jobs_config),
        migrations.RunPython(migrate_tensorboard_jobs_config),
        migrations.RunPython(migrate_experimentgroup_config),
        migrations.RunPython(create_cluster_owner),
    ]
