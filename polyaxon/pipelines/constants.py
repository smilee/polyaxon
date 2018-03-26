class OperationStatus(object):
    CREATED = 'created'
    SCHEDULED = 'scheduled'
    RUNNING = 'running'
    SUCCESS = 'success'
    FAILED = 'failed'
    UPSTREAM_FAILED = 'upstream_failed'
    STOPPED = 'stopped'
    SKIPPED = 'skipped'
    RETRYING = 'retrying'

    VALUES = [
        CREATED, SCHEDULED, RUNNING, SUCCESS, FAILED, UPSTREAM_FAILED, STOPPED, SKIPPED, RETRYING
    ]
    CHOICES = (
        (CREATED, CREATED),
        (SCHEDULED, SCHEDULED),
        (RUNNING, RUNNING),
        (SUCCESS, SUCCESS),
        (FAILED, FAILED),
        (UPSTREAM_FAILED, UPSTREAM_FAILED),
        (STOPPED, STOPPED),
        (SKIPPED, SKIPPED),
        (RETRYING, RETRYING)
    )

    DONE_STATUS = [SUCCESS, FAILED, UPSTREAM_FAILED, STOPPED, SKIPPED]

    ALLOWED_VALUES = {
        CREATED: set([]),
        SCHEDULED: {CREATED, RETRYING, },
        RUNNING: {SCHEDULED, },
        SUCCESS: {RUNNING, },
        FAILED: {SCHEDULED, RUNNING, },
        UPSTREAM_FAILED: set(VALUES),
        STOPPED: {CREATED, SCHEDULED, RUNNING, },
        SKIPPED: {CREATED, SCHEDULED, STOPPED, },
        RETRYING: {SCHEDULED, FAILED, STOPPED, SKIPPED},
    }

    @classmethod
    def can_transition(cls, status_from, status_to):
        return status_to in cls.ALLOWED_VALUES[status_from]


class TriggerRule(object):
    ALL_SUCCESS = 'all_success'
    ALL_FAILED = 'all_failed'
    ALL_DONE = 'all_done'
    ONE_SUCCESS = 'one_success'
    ONE_FAILED = 'one_failed'
    ONE_DONE = 'one_done'

    VALUES = [ALL_SUCCESS, ALL_FAILED, ALL_DONE, ONE_SUCCESS, ONE_FAILED, ONE_DONE]
    CHOICES = (
        (ALL_SUCCESS, ALL_SUCCESS),
        (ALL_FAILED, ALL_FAILED),
        (ALL_DONE, ALL_DONE),
        (ONE_SUCCESS, ONE_SUCCESS),
        (ONE_FAILED, ONE_FAILED),
        (ONE_DONE, ONE_DONE),
    )
