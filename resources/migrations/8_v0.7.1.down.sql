ALTER TABLE IF EXISTS "aws_config_configuration_recorders" DROP COLUMN status_last_error_code,
                                                           DROP COLUMN status_last_error_message,
                                                           DROP COLUMN status_last_start_time,
                                                           DROP COLUMN status_last_status,
                                                           DROP COLUMN status_last_status_change_time,
                                                           DROP COLUMN status_last_stop_time,
                                                           DROP COLUMN status_recording;
