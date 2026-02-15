CREATE TABLE attendances (
    id UUID PRIMARY KEY,
    employee_id UUID NOT NULL,
    check_in TIMESTAMP NOT NULL,
    check_out TIMESTAMP NULL,
    status SMALLINT,

    CONSTRAINT fk_employee
     FOREIGN KEY (employee_id)
         REFERENCES employees(id)
         ON DELETE CASCADE
);

CREATE UNIQUE INDEX unique_employee_daily
    ON attendances (employee_id, DATE(check_in));