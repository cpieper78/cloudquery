# Table: gitlab_groups_groups



The composite primary key for this table is (**id**, **name**).

## Relations

The following tables depend on gitlab_groups_groups:
  - [gitlab_groups_group_members](gitlab_groups_group_members.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|name (PK)|String|
|path|String|
|description|String|
|membership_lock|Bool|
|visibility|String|
|lfs_enabled|Bool|
|default_branch_protection|Int|
|avatar_url|String|
|web_url|String|
|request_access_enabled|Bool|
|full_name|String|
|full_path|String|
|file_template_project_id|Int|
|parent_id|Int|
|projects|JSON|
|statistics|JSON|
|custom_attributes|JSON|
|share_with_group_lock|Bool|
|require_two_factor_authentication|Bool|
|two_factor_grace_period|Int|
|project_creation_level|String|
|auto_devops_enabled|Bool|
|subgroup_creation_level|String|
|emails_disabled|Bool|
|mentions_disabled|Bool|
|runners_token|String|
|shared_projects|JSON|
|shared_runners_enabled|Bool|
|shared_with_groups|JSON|
|ldap_cn|String|
|ldap_access|Int|
|ldap_group_links|JSON|
|saml_group_links|JSON|
|shared_runners_minutes_limit|Int|
|extra_shared_runners_minutes_limit|Int|
|prevent_forking_outside_group|Bool|
|marked_for_deletion_on|JSON|
|created_at|Timestamp|