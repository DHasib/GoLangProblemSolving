# Rent-Management-System

This is a comprehensive tool named “RMS” with an aim of maximizing rental income and minimizing administrative overhead.The system will be designed to help landlords manage their rental properties more efficiently and effectively.

## ER Diagram

```mermaid
erDiagram
    Tenant ||--o{ Property : rents
    Tenant ||--o{ User: has
    User ||--o{ Renter : is
    User ||--o{ UserRole: has
    Role ||--o{ UserRole: has
    Property ||--o{ PropertyImage : has
    Property ||--|{ Renter : has
    Property ||--|{ PropertyUtility : has
    PropertyUtility ||--o{ UtilityReading : has
    User ||--|{ ConversationParticipent : has
    Conversation ||--|{ ConversationParticipent : has
    Conversation ||--|{ Message : has
    User ||--o{ Message: has
    Renter ||--|{ Rent : pays
    Renter ||--o{ Payment : "paid by"

    Tenant{
        id number PK
        name varchar
        code varchar
        details varchar
        is_active boolean
        created_at datetime
        updated_at datetime
    }

    User {
        id number PK
        tenant_id number FK
        username varchar
        email varchar
        password varchar
        is_superuser varchar
        created_at datetime
        updated_at datetime
    }

    Role {
        id number PK
        name varchar
        created_at datetime
        updated_at datetime
    }

    UserRole{
        id number PK
        user_id number FK
        role_id number FK
        created_at datetime
        updated_at datetime
    }

    Payment {
        id number PK
        user_id number FK
        tenant_id number FK
        payment_type varchar "cash, online"
        month number
        year number
        amount decimal
        created_at datetime
        updated_at datetime
    }

    Conversation {
        id number PK
        conversation_type varchar "private, Group"
        created_at datetime
        updated_at datetime
    }

    ConversationParticipent{
        id number PK
        conversation_id number FK
        participent_id number FK
    }

    Message{
        id number PK
        user_id number FK
        conversation_id number FK
        message_content varchar
        created_at datetime
        updated_at datetime

    }

    Renter{
        id number PK
        user_id number FK
        property_id number FK
        tenant_id number FK
        renat_amount decimal
        is_active boolean
        dues decimal
        starting_from datetime
        agreement_time datetime
        notice_period datetime
        created_at datetime
        updated_at datetime
    }

    Rent{
        id number PK
        renter_id number FK
        rent_amount deciaml
        year number
        month number
        start_from datetime
        is_calculated boolean
        created_at datetime
        updated_at datetime
    }

    Property{
        id number PK
        tenant_id number FK
        property_number varchar
        Property_details varchar
        property_rent_amount decimal
        property_room_count decimal
        property_square_feet number
        property_type varchar "residential, commercial"
        is_available boolean
        is_booked boolean
        added_by number FK
        created_at datetime
        updated_at datetime
    }

    PropertyImage{
        id number PK
        tenant_id number FK
        property_id number FK
        url string
        created_at datetime
        updated_at datetime
    }

    PropertyUtility{
        id number PK
        utility_number varchar
        utility_type varchar "water, gas, electricity"
        property_id number
        is_active boolean
        tenant_id number FK
        created_at datetime
        updated_at datetime
    }

    UtilityReading{
        id number PK
        month number
        year number
        reading decimal
        previous_reading_id number FK
        utility_id number FK
        consumed_unit decimal
        unit_price decimal
        month_bill decimal
        tenant_id number
        is_calculated boolean
        created_at datetime
        updated_at datetime
    }

    BlacklistedToken{
        id number PK
        token varchar
        created_at datetime
        updated_at datetime
    }


```

## Activity Diagrams

```mermaid
---
---
title: Update Renter Balance on Reading Input
---
flowchart LR
    start((START))
    stop(((STOP)))

    start --> A
    A[/Get utility/] --> B{Utility Exists ?}
    B --No--> Exception[Rise Exception]
    B --Yes--> C{Same month reading\n exists for same utility ?}
    C --Yes--> Exception
    C --No--> D[/Get active Renter of the property/]
    D --No--> E{Multiple active renter ?}
    E --Yes--> Exception
    E --No--> F[/Store reading/]
    F --> G[/Update renter's balance/]
    G --> stop
```

```mermaid
---
---
title: Rent Payment
---
flowchart LR
    start((Start))
    stop(((STOP)))

    Exception[/Raise Exception/]
    start --> A
    A[/Get Renter/] --> B{Renter exists ?}
    B --No--> Exception
    B --Yes--> C{Is pay on cash?}
    C --Yes--> D[/store payment data/]
    D --> E[/Update Renter Balance/]
    C --No--> F[Get Payment Gateway's\n response]
    F --> D
    E --> stop
```
