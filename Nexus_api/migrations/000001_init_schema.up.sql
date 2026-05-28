DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'goal_status') THEN
        CREATE TYPE goal_status AS ENUM (
            'Pendente',
            'Concluido',
            'Cancelado',
            'Atrasada'
        );
    END IF;
END;
$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
        CREATE TYPE task_status AS ENUM (
            'Pendente',
            'Em progresso',
            'Concluido'
        );
    END IF;
END;
$$;

CREATE TABLE IF NOT EXISTS users (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    public_id uuid NOT NULL,
    name varchar(150) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(11) NOT NULL,
    password varchar(8) NOT NULL,
    CONSTRAINT uni_users_public_id UNIQUE (public_id),
    CONSTRAINT uni_users_email UNIQUE (email)
);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at ON users (deleted_at);

CREATE TABLE IF NOT EXISTS goals (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(150) NOT NULL,
    description text DEFAULT 'Descrição não informada',
    start_date timestamptz NOT NULL,
    finalization_forecast timestamptz,
    status goal_status NOT NULL DEFAULT 'Pendente',
    user_id bigint NOT NULL,
    CONSTRAINT fk_users_goals FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX IF NOT EXISTS idx_goals_deleted_at ON goals (deleted_at);
CREATE INDEX IF NOT EXISTS idx_goals_user_id ON goals (user_id);

CREATE TABLE IF NOT EXISTS tasks (
    id bigserial PRIMARY KEY,
    created_at timestamptz,
    updated_at timestamptz,
    deleted_at timestamptz,
    name varchar(150) NOT NULL,
    description text DEFAULT 'Descrição não informada',
    status task_status NOT NULL DEFAULT 'Pendente',
    start_date timestamptz,
    finalization_date timestamptz,
    time_spent bigint,
    goal_id bigint NOT NULL,
    CONSTRAINT fk_goals_tasks FOREIGN KEY (goal_id) REFERENCES goals (id)
);

CREATE INDEX IF NOT EXISTS idx_tasks_deleted_at ON tasks (deleted_at);
CREATE INDEX IF NOT EXISTS idx_tasks_goal_id ON tasks (goal_id);
