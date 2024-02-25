if [ ${ENV} = "DEV" ]; then 
   pnpm run dev
else
   node .output/server/index.mjs
fi
