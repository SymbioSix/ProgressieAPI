# Progressie API

RESTful Progressie Academy API Services. Built to ensure Progressie Services are good to be served!

## Development Inquiry

Simply fork this repository. Then follow these steps :

```bash
  1. cmd: cd ProgressieAPI || cd <your-project>
  2. cmd: go get -u
  3. copy the .env.example file and paste + rename it into .env
  4. open the .env and fill the corresponding environment variables (the values will be shared only among owners)
  5. do your assigned job properly
  6. do not forget to add API Documentation comments in your service file (example: services/auth/index.go)
  7. cmd: swag fmt
  8. cmd: swag init
  9. make sure the swag has no error on parsing your API Docs comments (if error, pls check Swagger Docs in References section for fix)
  10. cmd: go run main.go
  11. make sure that firewall allow access popped up + allowed and no errors/conflicts (bugs are tolerated, but beware).
  12. cmd: git checkout -b [YOUR_BRANCH]
  13. cmd: git commit -am '[YOUR COMMIT MESSAGES]'
  14. cmd: git push origin [YOUR_BRANCH]
  15. Create New Pull Request
```

## References

- [GoFiber](https://docs.gofiber.io/next/)
- [Supabase Community](https://github.com/supabase-community)
- [Previous Private Repository](https://github.com/krisnaganesha1609/smapurv1_api)
- [Swagger Docs](https://github.com/swaggo/swag?tab=readme-ov-file#declarative-comments-format)
