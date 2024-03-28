env "local" {
  src = [
    "file://schema.my.hcl",
  ]
  dev = "mysql://root:root@localhost:3306/test"
  url = "mysql://root:root@localhost:3306/auth_db"
  migration {
    dir = "file://migrations"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
