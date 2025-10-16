# ApiService - User Management

## GetUserProfile

Retrieves information about the currently authenticated user.

```proto
GetUserProfile() returns (io.clbs.openhes.models.system.UserProfile)
```

- Output: [`io.clbs.openhes.models.system.UserProfile`](model-io-clbs-openhes-models-system-userprofile.md)

## UpdateUserProfile

Updates the profile of the user identified by the UserProfile's id field. Read only fields will be ignored. Permissions may apply to update other user's profiles.

```proto
UpdateUserProfile(io.clbs.openhes.models.system.UserProfile)
```

- Input: [`io.clbs.openhes.models.system.UserProfile`](model-io-clbs-openhes-models-system-userprofile.md)

