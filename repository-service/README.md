# Repository Service

Create Secret 
berglas create github-ntf-secrets/NAME SECRET --key projects/${PROJECT_ID}/locations/global/keyRings/berglas/cryptoKeys/berglas-key
berglas grant github-ntf-secrets/mongourl --member "serviceAccount:service-862979686689@gcf-admin-robot.iam.gserviceaccount.com"