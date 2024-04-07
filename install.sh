go install github.com/mitranim/gow@latest
cd ./apps/api && go mod tidy

npm i -g pnpm turbo
pnpm i
turbo prepack

if [ ${ENV} = "PROD" ]; then 
   turbo build
fi

