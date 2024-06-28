table "cookies" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "cookie" {
    null = false
    type = text
  }
  column "suffix" {
    null = false
    type = varchar(16)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "cookies_created_index" {
    columns = [column.created]
  }
  index "cookies_deleted_index" {
    columns = [column.deleted]
  }
  index "cookies_modified_index" {
    columns = [column.modified]
  }
}
table "deliveries_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_50mais_created_index" {
    columns = [column.created]
  }
  index "deliveries_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_50mais_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_50mais_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_50mais_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_50mais_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_50mais_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_coomap_created_index" {
    columns = [column.created]
  }
  index "deliveries_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_coomap_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_coomap_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_coomap_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_coomap_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_coomap_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "deliveries_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_ecoexpress_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_ecoexpress_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_ecoexpress_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_ecoexpress_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_ecoexpress_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_elologistica_created_index" {
    columns = [column.created]
  }
  index "deliveries_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_elologistica_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_elologistica_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_elologistica_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_elologistica_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_elologistica_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_mpl_created_index" {
    columns = [column.created]
  }
  index "deliveries_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_mpl_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_mpl_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_mpl_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_mpl_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_mpl_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_parceiro_created_index" {
    columns = [column.created]
  }
  index "deliveries_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_parceiro_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_parceiro_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_parceiro_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_parceiro_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_parceiro_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "deliveries_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "entrega_id" {
    null = true
    type = bigint
  }
  column "rota_id" {
    null = true
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "parada_id" {
    null = true
    type = char(26)
  }
  column "address" {
    null = true
    type = varchar(255)
  }
  column "addressType" {
    null = true
    type = varchar(255)
  }
  column "locationLat" {
    null = true
    type = varchar(255)
  }
  column "locationLng" {
    null = true
    type = varchar(255)
  }
  column "horaEntregue" {
    null = true
    type = datetime
  }
  column "numero_de_rota_planejada" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "subStatus" {
    null = true
    type = varchar(255)
  }
  column "ordem" {
    null = true
    type = int
  }
  column "fora_de_area" {
    null = true
    type = bool
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "deliveries_rodacoop_created_index" {
    columns = [column.created]
  }
  index "deliveries_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "deliveries_rodacoop_horaentregue_index" {
    columns = [column.horaEntregue]
  }
  index "deliveries_rodacoop_modified_index" {
    columns = [column.modified]
  }
  index "deliveries_rodacoop_parada_id_index" {
    columns = [column.parada_id]
  }
  index "deliveries_rodacoop_rota_id_index" {
    columns = [column.rota_id]
  }
  index "deliveries_rodacoop_sc_id_index" {
    columns = [column.sc_id]
  }
}
table "drivers_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_50mais_created_index" {
    columns = [column.created]
  }
  index "drivers_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_50mais_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_50mais_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_50mais_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_50mais_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_coomap_created_index" {
    columns = [column.created]
  }
  index "drivers_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_coomap_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_coomap_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_coomap_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_coomap_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "drivers_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_ecoexpress_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_ecoexpress_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_ecoexpress_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_ecoexpress_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_elologistica_created_index" {
    columns = [column.created]
  }
  index "drivers_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_elologistica_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_elologistica_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_elologistica_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_elologistica_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_mpl_created_index" {
    columns = [column.created]
  }
  index "drivers_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_mpl_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_mpl_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_mpl_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_mpl_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_parceiro_created_index" {
    columns = [column.created]
  }
  index "drivers_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_parceiro_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_parceiro_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_parceiro_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_parceiro_modified_index" {
    columns = [column.modified]
  }
}
table "drivers_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "driverSystemId" {
    null = false
    type = int
  }
  column "carrierId" {
    null = false
    type = int
  }
  column "identificationType" {
    null = false
    type = varchar(32)
  }
  column "identificationValue" {
    null = false
    type = varchar(32)
  }
  column "name" {
    null = false
    type = varchar(128)
  }
  column "firstName" {
    null = false
    type = varchar(64)
  }
  column "lastName" {
    null = false
    type = varchar(64)
  }
  column "email" {
    null = true
    type = varchar(255)
  }
  column "phone" {
    null = false
    type = varchar(255)
  }
  column "creationDate" {
    null = false
    type = int
  }
  column "lastUpdated" {
    null = false
    type = int
  }
  column "status" {
    null = false
    type = varchar(255)
  }
  column "disabled" {
    null = false
    type = bool
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "drivers_rodacoop_created_index" {
    columns = [column.created]
  }
  index "drivers_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "drivers_rodacoop_disabled_index" {
    columns = [column.disabled]
  }
  index "drivers_rodacoop_driverSystemId_index" {
    unique  = true
    columns = [column.driverSystemId]
  }
  index "drivers_rodacoop_identificationvalue_index" {
    columns = [column.identificationValue]
  }
  index "drivers_rodacoop_modified_index" {
    columns = [column.modified]
  }
}
table "migrations" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null           = false
    type           = int
    unsigned       = true
    auto_increment = true
  }
  column "migration" {
    null = false
    type = varchar(255)
  }
  column "batch" {
    null = false
    type = int
  }
  primary_key {
    columns = [column.id]
  }
}
table "motoristas_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_50mais_created_index" {
    columns = [column.created]
  }
  index "motoristas_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_50mais_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_50mais_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_50mais_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_coomap_created_index" {
    columns = [column.created]
  }
  index "motoristas_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_coomap_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_coomap_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_coomap_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "motoristas_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_ecoexpress_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_ecoexpress_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_ecoexpress_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_elologistica_created_index" {
    columns = [column.created]
  }
  index "motoristas_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_elologistica_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_elologistica_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_elologistica_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_mpl_created_index" {
    columns = [column.created]
  }
  index "motoristas_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_mpl_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_mpl_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_mpl_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_parceiro_created_index" {
    columns = [column.created]
  }
  index "motoristas_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_parceiro_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_parceiro_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_parceiro_sc_index" {
    columns = [column.sc]
  }
}
table "motoristas_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "meliId" {
    null = false
    type = int
  }
  column "experience" {
    null = false
    type = varchar(255)
  }
  column "driverStatus" {
    null = false
    type = varchar(255)
  }
  column "sc" {
    null = false
    type = varchar(255)
  }
  column "contactRate" {
    null     = false
    type     = double(8)
    unsigned = false
  }
  column "claimsCount" {
    null = false
    type = int
  }
  column "shipmentsCount" {
    null = false
    type = int
  }
  column "missingShipmentsCount" {
    null = false
    type = int
  }
  column "stolenShipmentsCount" {
    null = false
    type = int
  }
  column "lostShipmentsCount" {
    null = false
    type = int
  }
  column "emptyBoxCount" {
    null = false
    type = int
  }
  column "blockingSoon" {
    null = false
    type = bool
  }
  column "lastRouteDate" {
    null = false
    type = datetime
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "motoristas_rodacoop_created_index" {
    columns = [column.created]
  }
  index "motoristas_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "motoristas_rodacoop_meliid_index" {
    columns = [column.meliId]
  }
  index "motoristas_rodacoop_modified_index" {
    columns = [column.modified]
  }
  index "motoristas_rodacoop_sc_index" {
    columns = [column.sc]
  }
}
table "routes_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_50mais_created_index" {
    columns = [column.created]
  }
  index "routes_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_50mais_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_50mais_largada_index" {
    columns = [column.largada]
  }
  index "routes_50mais_modified_index" {
    columns = [column.modified]
  }
  index "routes_50mais_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_coomap_created_index" {
    columns = [column.created]
  }
  index "routes_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_coomap_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_coomap_largada_index" {
    columns = [column.largada]
  }
  index "routes_coomap_modified_index" {
    columns = [column.modified]
  }
  index "routes_coomap_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "routes_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_ecoexpress_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_ecoexpress_largada_index" {
    columns = [column.largada]
  }
  index "routes_ecoexpress_modified_index" {
    columns = [column.modified]
  }
  index "routes_ecoexpress_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_elologistica_created_index" {
    columns = [column.created]
  }
  index "routes_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_elologistica_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_elologistica_largada_index" {
    columns = [column.largada]
  }
  index "routes_elologistica_modified_index" {
    columns = [column.modified]
  }
  index "routes_elologistica_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_mpl_created_index" {
    columns = [column.created]
  }
  index "routes_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_mpl_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_mpl_largada_index" {
    columns = [column.largada]
  }
  index "routes_mpl_modified_index" {
    columns = [column.modified]
  }
  index "routes_mpl_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_parceiro_created_index" {
    columns = [column.created]
  }
  index "routes_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_parceiro_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_parceiro_largada_index" {
    columns = [column.largada]
  }
  index "routes_parceiro_modified_index" {
    columns = [column.modified]
  }
  index "routes_parceiro_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "routes_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "id_old" {
    null = true
    type = int
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "sc_id" {
    null = true
    type = char(26)
  }
  column "facility" {
    null = true
    type = varchar(255)
  }
  column "cluster" {
    null = true
    type = varchar(255)
  }
  column "placa" {
    null = true
    type = varchar(255)
  }
  column "motorista" {
    null = true
    type = varchar(255)
  }
  column "exp_motorista" {
    null = true
    type = varchar(255)
  }
  column "progresso" {
    null = true
    type = varchar(255)
  }
  column "entregues" {
    null = true
    type = int
  }
  column "pendentes" {
    null = true
    type = int
  }
  column "falhas" {
    null = true
    type = int
  }
  column "spr" {
    null = true
    type = int
  }
  column "entregue_fora_de_area" {
    null = true
    type = int
  }
  column "falha_entregue_fora_de_area" {
    null = true
    type = int
  }
  column "orh" {
    null = true
    type = int
  }
  column "ozh" {
    null = true
    type = int
  }
  column "stem_in" {
    null = true
    type = int
  }
  column "stem_out" {
    null = true
    type = int
  }
  column "rota_estimada" {
    null = true
    type = int
  }
  column "traveled_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "planned_distance" {
    null     = true
    type     = double(8)
    unsigned = false
  }
  column "has_helper" {
    null = true
    type = bool
  }
  column "reclamcao_aberta" {
    null = true
    type = int
  }
  column "reclamcao_fechada" {
    null = true
    type = int
  }
  column "veiculo" {
    null = true
    type = varchar(255)
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "inicioRota" {
    null = true
    type = timestamp
  }
  column "largada" {
    null = true
    type = timestamp
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "routes_rodacoop_created_index" {
    columns = [column.created]
  }
  index "routes_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "routes_rodacoop_id_old_unique" {
    unique  = true
    columns = [column.id_old]
  }
  index "routes_rodacoop_largada_index" {
    columns = [column.largada]
  }
  index "routes_rodacoop_modified_index" {
    columns = [column.modified]
  }
  index "routes_rodacoop_rota_id_index" {
    columns = [column.rota_id]
  }
}
table "stops_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_50mais_created_index" {
    columns = [column.created]
  }
  index "stops_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_50mais_modified_index" {
    columns = [column.modified]
  }
  index "stops_50mais_parada_id_index" {
    columns = [column.parada_id]
  }
  index "stops_50mais_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_50mais_status_index" {
    columns = [column.status]
  }
}
table "stops_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_coomap_created_index" {
    columns = [column.created]
  }
  index "stops_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_coomap_modified_index" {
    columns = [column.modified]
  }
  index "stops_coomap_parada_id_index" {
    columns = [column.parada_id]
  }
  index "stops_coomap_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_coomap_status_index" {
    columns = [column.status]
  }
}
table "stops_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "stops_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_ecoexpress_modified_index" {
    columns = [column.modified]
  }
  index "stops_ecoexpress_parada_id_index" {
    columns = [column.parada_id]
  }
  index "stops_ecoexpress_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_ecoexpress_status_index" {
    columns = [column.status]
  }
}
table "stops_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_elologistica_created_index" {
    columns = [column.created]
  }
  index "stops_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_elologistica_modified_index" {
    columns = [column.modified]
  }
  index "stops_elologistica_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_elologistica_status_index" {
    columns = [column.status]
  }
}
table "stops_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_mpl_created_index" {
    columns = [column.created]
  }
  index "stops_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_mpl_modified_index" {
    columns = [column.modified]
  }
  index "stops_mpl_parada_id_index" {
    columns = [column.parada_id]
  }
  index "stops_mpl_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_mpl_status_index" {
    columns = [column.status]
  }
}
table "stops_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_parceiro_created_index" {
    columns = [column.created]
  }
  index "stops_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_parceiro_modified_index" {
    columns = [column.modified]
  }
  index "stops_parceiro_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_parceiro_status_index" {
    columns = [column.status]
  }
}
table "stops_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "parada_id" {
    null = false
    type = bigint
  }
  column "rota_id" {
    null = false
    type = bigint
  }
  column "status" {
    null = true
    type = varchar(255)
  }
  column "endereco" {
    null = true
    type = varchar(255)
  }
  column "receivedRouteId" {
    null = true
    type = varchar(255)
  }
  column "transferredRouteId" {
    null = true
    type = varchar(255)
  }
  column "preferenciaEntrega" {
    null = true
    type = varchar(255)
  }
  column "quantidadePedidos" {
    null = true
    type = int
  }
  column "quantidadeUnidadeTransportes" {
    null = true
    type = int
  }
  column "quantidadePedidosAfetados" {
    null = true
    type = int
  }
  column "tocTotalCases" {
    null = true
    type = int
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "stops_rodacoop_created_index" {
    columns = [column.created]
  }
  index "stops_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "stops_rodacoop_modified_index" {
    columns = [column.modified]
  }
  index "stops_rodacoop_parada_id_uindex" {
    unique  = true
    columns = [column.parada_id]
  }
  index "stops_rodacoop_paradas_id_index" {
    columns = [column.parada_id]
  }
  index "stops_rodacoop_rota_id_index" {
    columns = [column.rota_id]
  }
  index "stops_rodacoop_status_index" {
    columns = [column.status]
  }
}
table "system_configs" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "config" {
    null = false
    type = varchar(100)
  }
  column "value" {
    null = false
    type = text
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "system_configs_config_index" {
    columns = [column.config]
  }
  index "system_configs_created_index" {
    columns = [column.created]
  }
  index "system_configs_deleted_index" {
    columns = [column.deleted]
  }
  index "system_configs_modified_index" {
    columns = [column.modified]
  }
}
table "xpts_50mais" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_50mais_created_index" {
    columns = [column.created]
  }
  index "xpts_50mais_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_50mais_modified_index" {
    columns = [column.modified]
  }
  index "xpts_50mais_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_coomap" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_coomap_created_index" {
    columns = [column.created]
  }
  index "xpts_coomap_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_coomap_modified_index" {
    columns = [column.modified]
  }
  index "xpts_coomap_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_ecoexpress" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_ecoexpress_created_index" {
    columns = [column.created]
  }
  index "xpts_ecoexpress_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_ecoexpress_modified_index" {
    columns = [column.modified]
  }
  index "xpts_ecoexpress_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_elologistica" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_elologistica_created_index" {
    columns = [column.created]
  }
  index "xpts_elologistica_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_elologistica_modified_index" {
    columns = [column.modified]
  }
  index "xpts_elologistica_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_mpl" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_mpl_created_index" {
    columns = [column.created]
  }
  index "xpts_mpl_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_mpl_modified_index" {
    columns = [column.modified]
  }
  index "xpts_mpl_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_parceiro" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_parceiro_created_index" {
    columns = [column.created]
  }
  index "xpts_parceiro_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_parceiro_modified_index" {
    columns = [column.modified]
  }
  index "xpts_parceiro_sigla_index" {
    columns = [column.sigla]
  }
}
table "xpts_rodacoop" {
  schema  = schema.backend-v2
  collate = "utf8mb4_unicode_ci"
  column "id" {
    null = false
    type = char(26)
  }
  column "sigla" {
    null = false
    type = varchar(255)
  }
  column "lat" {
    null = false
    type = varchar(255)
  }
  column "lng" {
    null = false
    type = varchar(255)
  }
  column "created" {
    null = false
    type = datetime
  }
  column "modified" {
    null = false
    type = datetime
  }
  column "deleted" {
    null = true
    type = datetime
  }
  primary_key {
    columns = [column.id]
  }
  index "xpts_rodacoop_created_index" {
    columns = [column.created]
  }
  index "xpts_rodacoop_deleted_index" {
    columns = [column.deleted]
  }
  index "xpts_rodacoop_modified_index" {
    columns = [column.modified]
  }
  index "xpts_rodacoop_sigla_index" {
    columns = [column.sigla]
  }
}
schema "backend-v2" {
  charset = "utf8mb4"
  collate = "utf8mb4_0900_ai_ci"
}
