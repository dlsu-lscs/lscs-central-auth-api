# LSCS Central Auth API

The official *Authentication Microservice* of **La Salle Computer Society (LSCS)**

This is an auth microservice, meant to be used by an application backend.

_**Treat this as a service that simply returns a JSON payload, used only for authenticating LSCS Members and returning necessary data from them.**_

## Usage

> [!IMPORTANT]
> **only RND members can request an API key (associated with their DLSU email)** - to prevent unauthorized access


## Auth Endpoints

- `[UPDATE 20250107-04:49AM]`: **no longer needs to login to Google**

### POST `/request-key`

- `request`:
```bash
curl -X POST http://localhost:42069/member \
  -H "Content-Type: application/json" \
  -d '{"email": "edwin_sadiarinjr@dlsu.edu.ph"}'
```

- `response`:
```json
{
    "api_key": "somethingsomethingstringssdgasdfkgjdsf",
    "email": "edwin_sadiarinjr@dlsu.edu.ph"
}
```

### POST `/revoke-key`

- for revoking/deleting key

- `request`:
```bash
curl -X POST http://localhost:42069/request-key \
  -H "Content-Type: application/json" \
  -d '{"email": "edwin_sadiarinjr@dlsu.edu.ph"}'
  -d '{"pepper": "<CONTACT_ADMIN_DEVELOPER_TO_REVOKE_KEY>"}'
```

- `response`:
```
API key for <email> is successfully revoked
```


## Member Endpoints

### GET `/members`

- returns all LSCS members from database (*yes*)

- `request`:
```bash
curl -X GET http://localhost:42069/member \
  -H "Authorization: Bearer <API-KEY>"
```

- `response`:
```json
[
   {
     "id": 12312312,
     "full_name": "Hehe E. Hihi",
     "nickname": "Huhi",
     "email": "hehe_hihi@dlsu.edu.ph",
     "telegram": "",
     "position_id": "MEM",
     "committee_id": "MEM",
     "college": "CCS",
     "program": "BS-Org",
     "discord": ""
   },
   {
     "id": 11111110,
     "full_name": "Peter Parker",
     "nickname": "Peter",
     "email": "peter_parker@dlsu.edu.ph",
     "telegram": "@something",
     "position_id": "MEM",
     "committee_id": "MEM",
     "college": "CLA",
     "program": "POM-LGL",
     "discord": ""
   }
]
```

### POST `/member`

- returns `email`, `full_name`, `committee_name`, `position_name`, and `division_name` of the LSCS member 

- `request`:
```bash
curl -X POST http://localhost:42069/member \
  -H "Authorization: Bearer <API-KEY>" \
  -H "Content-Type: application/json" \
  -d '{"email": "edwin_sadiarinjr@dlsu.edu.ph"}'
```

- `response`:
```json
{ // success
  "committee_name": "Research and Development",
  "division_name": "Internals",
  "email": "edwin_sadiarinjr@dlsu.edu.ph",
  "full_name": "Edwin Sadiarin Jr.",
  "position_name": "Committee Trainee"
}

{ // fail
  "error": "Email is not an LSCS member"
}
```

### POST `/check-email`

- checks if the email exists in database (indicating if it is an LSCS member or not)

- `request`:
```bash
curl -X POST http://localhost:42069/check-email \
  -H "Authorization: Bearer <API-KEY>" \
  -H "Content-Type: application/json" \
  -d '{"email": "edwin_sadiarinjr@dlsu.edu.ph"}'
```

- `response`:
```json
{ // success
  "email": "edwin_sadiarinjr@dlsu.edu.ph",
  "state": "present",
  "success": "Email is an LSCS member"
}

{ // fail
  "email": "test@dlsu.edu.ph",
  "error": "Not an LSCS member",
  "state": "absent"
}
```
