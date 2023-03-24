schema "public" {}

table "todos" {
  schema = schema.public

  column "id" {
    null = false
    type = character_varying(36)
  }
  column "name" {
    null = false
    type = character_varying(50)
  }
  column "description" {
    null = true
    type = text
  }
  column "status" {
    null = true
    type = text
  }
  column "createdAt" {
    null    = false
    type    = timestamp(3)
    default = sql("CURRENT_TIMESTAMP")
  }
  column "updatedAt" {
    null = false
    type = timestamp(3)
  }

  primary_key {
    columns = [column.id]
  }
}