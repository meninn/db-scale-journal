-- Schema inicial (Fase 0) — v0 ingênua DE PROPÓSITO.
--
-- Regras desta versão:
--   * Cada tabela tem apenas uma PK surrogate (id AUTO_INCREMENT).
--   * Nenhum índice secundário, nenhum UNIQUE, nenhuma FOREIGN KEY.
--     (No InnoDB, FK cria índice automaticamente na coluna referenciadora,
--      o que já "resolveria" parte dos gargalos antes de medi-los.)
--   * Integridade (usuário existe, follow duplicado, seguir a si mesmo)
--     fica a cargo da aplicação na v0.
--
-- Índices, chaves compostas, constraints e desnormalização são exatamente
-- o que o ciclo da Fase 5+ deve descobrir e justificar com medição.

CREATE TABLE users (
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    username   VARCHAR(50)     NOT NULL,
    created_at DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE posts (
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id    BIGINT UNSIGNED NOT NULL,
    body       VARCHAR(280)    NOT NULL,
    created_at DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- follower_id segue followee_id.
CREATE TABLE follows (
    id          BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    follower_id BIGINT UNSIGNED NOT NULL,
    followee_id BIGINT UNSIGNED NOT NULL,
    created_at  DATETIME(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
