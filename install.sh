go install github.com/mitranim/gow@latest
npm i -g pnpm turbo
cd ./apps/api && go mod tidy
pnpm i
turbo prepack

if [ ${ENV} = "PROD" ]; then 
   turbo build
fi

