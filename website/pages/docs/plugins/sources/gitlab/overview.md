# Vercel Source Plugin

import { getLatestVersion } from "../../../../../utils/versions";
import { Badge } from "../../../../../components/Badge";

<Badge text={"Latest: " + getLatestVersion("source", `gitlab`)}/>

The CloudQuery GitLab plugin pulls configuration out of GitLab resources and loads it into any supported CloudQuery destination (e.g. PostgreSQL).

## Authentication

In order to fetch information from GitLab, `cloudquery` needs to be authenticated. An access token is required for authentication. Instructions on how to generate an access token [here](https://docs.gitlab.com/ee/user/profile/personal_access_tokens.html#create-a-personal-access-token). Note that this token should only have read-only access to the resources you intend to use.


## Query Examples

TODO