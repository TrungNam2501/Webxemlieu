# Webxemlieu

Web xem dữ liệu của công ty cao su Kenda. Dự án gồm 2 phần:

- `backend/` — REST API viết bằng **Go (Gin + GORM/PostgreSQL)**.
- `frontend/` — SPA viết bằng **Vue 3 + Vite + Tailwind CSS v4**.

Trong production, backend serve luôn build của frontend dưới prefix `/app/`.

## Yêu cầu

| Công cụ    | Phiên bản |
| ---------- | --------- |
| Go         | ≥ 1.24    |
| Node.js    | ≥ 20      |
| PostgreSQL | có quyền truy cập database `kverp` |

## 1. Backend

```bash
cd backend
cp .env.example .env             # chỉnh DB_HOST, DB_USER, DB_PASSWORD, …
go mod tidy
go run .
```

Mặc định backend lắng nghe trên `:8080`. Toàn bộ cấu hình (port, gin mode,
DB, CORS) đều đọc từ env, xem `backend/.env.example`.

### API

- `GET /api/items`
- `GET /api/recipe-definitions` — trả tối đa 100 dòng đầu của
  `kvmes.recipe_process_definition`.

### Phục vụ frontend (production)

Sau khi build frontend (`npm run build` trong `frontend/`), backend sẽ tự
serve `frontend/dist` ở `/app/`. Đường dẫn dist có thể đổi qua biến
`FRONTEND_DIST`.

## 2. Frontend

```bash
cd frontend
cp .env.example .env             # tuỳ chọn, để cấu hình proxy
npm install
npm run dev                      # http://localhost:5173/app/
```

Vite được cấu hình proxy `/api` → `http://localhost:8080` (đổi bằng
`VITE_API_PROXY_TARGET`). Khi chạy production, axios dùng đường dẫn
tương đối `/api`, nên FE và BE phải cùng origin.

### Lệnh khác

```bash
npm run lint        # eslint --fix
npm run type-check  # vue-tsc
npm run build       # type-check + vite build (output ra dist/)
npm run format      # prettier
```

## Cấu trúc thư mục

```
backend/
  config/        # đọc env, build DSN
  controllers/   # handlers Gin
  database/      # khởi tạo GORM
  middleware/    # CORS
  models/        # các struct GORM
  routes/        # đăng ký /api và SPA fallback
  main.go        # entrypoint, graceful shutdown
frontend/
  src/
    api/         # axios instance dùng chung
    components/  # ItemList.vue
    router/      # vue-router
    types/       # interface TypeScript
    views/       # RecipesView.vue
```
