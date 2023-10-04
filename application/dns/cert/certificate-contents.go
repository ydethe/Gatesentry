package dnsCerts

var CaPEM = []byte(`
-----BEGIN CERTIFICATE-----
MIIDeDCCAmCgAwIBAgIEZrIKfDANBgkqhkiG9w0BAQsFADBVMQswCQYDVQQGEwJO
TzETMBEGA1UEAwwKR2F0ZXNlbnRyeTENMAsGA1UECAwET3NsbzENMAsGA1UEBwwE
T3NsbzETMBEGA1UECgwKR2F0ZXNlbnRyeTAeFw0yMzEwMDQwOTA0MTRaFw0zMzEw
MDEwOTA0MTRaMFUxCzAJBgNVBAYTAk5PMRMwEQYDVQQDDApHYXRlc2VudHJ5MQ0w
CwYDVQQIDARPc2xvMQ0wCwYDVQQHDARPc2xvMRMwEQYDVQQKDApHYXRlc2VudHJ5
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEArhNHIfQpLKi+hBgT8kgt
LCITRklNwLuXlH2hsLEyTGIZMhx6BMdLOh/IgXhzjGL6TyLxgDrHAYqhs5p0kN7Y
7NSrnQXPrLKINEria0JjpkxNfyG1WswD0feRrT+iwM0XKBF0dttjXUwHoQ82uJgH
0P9PagP8V3I1bBQWPCNYjkSU7KCVuugM0eRVOsPK5tkK/uIwTAFHyehbNHNNI+6R
8RfU17mtSusV8gdSrVhaZBjLY5Cgm1ELD+a3cmKU1UiRBvFsnnF+CawKz6YSlk8/
T1ET2OhsWmsbTO3VgpM+TVkDTfIMHLf8YZ+jv+rvPg4VZti9/abwiM6ulxYu1yF/
fQIDAQABo1AwTjAdBgNVHQ4EFgQUke9IZER/F+s7+Nug0EmAkIBfmw8wHwYDVR0j
BBgwFoAUke9IZER/F+s7+Nug0EmAkIBfmw8wDAYDVR0TBAUwAwEB/zANBgkqhkiG
9w0BAQsFAAOCAQEAVft2SAxicaqBM/5Ek/AnkXVVs8jmZYI8wiZL69Fsdv8FZ1PO
uOF+CAY5+raUgiexRPbNPRDTyn8Tp8KcS9kxB2WpZSuV+kgP1ZsDmn5iwO6cVA5w
mmCQPEOa+kPhrJwMQd9kOMKz0Vl36lAGyLxxn2GpCRrIZO5kcflDQZsBvx7ejZDS
ohbp4BRhG2OWRw6LDMNYApwfvFmHyzlvYNEeXjvodk6UplP55KS7+gGPmS9IVzEi
IVOa35wFaLxDfoMJUgw93PsXihIqdMhRKhT3jfUEKhNsM1U3wJLyZ+zN1fAkIpRB
6AYb6nc22KyWwpBze+mzSO6sFfTwfXHO3GOPLA==
-----END CERTIFICATE-----
`)

var CaPrivKeyPEM = []byte(`
-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCuE0ch9CksqL6E
GBPySC0sIhNGSU3Au5eUfaGwsTJMYhkyHHoEx0s6H8iBeHOMYvpPIvGAOscBiqGz
mnSQ3tjs1KudBc+ssog0SuJrQmOmTE1/IbVazAPR95GtP6LAzRcoEXR222NdTAeh
Dza4mAfQ/09qA/xXcjVsFBY8I1iORJTsoJW66AzR5FU6w8rm2Qr+4jBMAUfJ6Fs0
c00j7pHxF9TXua1K6xXyB1KtWFpkGMtjkKCbUQsP5rdyYpTVSJEG8WyecX4JrArP
phKWTz9PURPY6GxaaxtM7dWCkz5NWQNN8gwct/xhn6O/6u8+DhVm2L39pvCIzq6X
Fi7XIX99AgMBAAECggEAO4uYn4K3uvEWCnV6DTynRxt44Gge0rjYxxCaiKh0fjBo
Uf3vt0a88usAWVlsnS8WVI+tcKGqhVp4qclB6nRwW8L72UCto3OGp+yduvcAw1gC
gmRzdtWm0OIQ0OSdtbnyG+CsqCOvv7BMQ8nCfC51LgbHkYo/fWIx6ACPAo9MsY0v
vFVf5cY0PdR9VGCXeezPBvo2Be9BMhZG+M6ymcMQaBUJLaZvC27PLd3uir4JYaT2
sI3KSzKjObpTXFR11sp15hR4PlmN9aadEN2LLXaPT2IobKbkpSkj+uS4PE+vxKYw
6VMKtKPCiQygRyzMtvTEn+ftlp4e3ToiwAgiESXzwQKBgQDWErtK3jsXgdIRw+LD
yn2/arGuITq1X7kwItqMNbB7CHboU0ZoKYbwWjbXL8wmqq+lDmD6j/KZRjnWZw6n
bdkPriZCix3zwE76eQzcy7KTVFtKjYm09+aFzMmCJuWzg6Y0mFwmNXd9rgCsR2fa
u7nHR2SnX6i+cyTNjYGznld/+QKBgQDQKx7dyy4Q1UgaqhWgjwblIGH3QolIVVDM
NPR1SIcxzLnZdkWVo4FldyRpB2s4rxVrPi6WnfAJdL4zh80hggQ5PRYA3W3Y0L7O
P68oX6Naz9i29+zCQWdzlaXmZ8KwoL+9PkDecE1c0o3kFTgxZyg1T+8UvPc2f/1V
DnrySeUkpQKBgDnDjC5DkQZizWBlzwat2QiRragi50iRr9LBVN+IjTplqlA+SD1L
F1I7xZiDGT9Fx2duXdS+iuO5A1pLLLUY+v00LLa/+zEOr0D+8O2TOXhvxsJzNrlS
Oy3XeHhaLpkl6O9APX1B6CBNl3jlO6zWAuc26H4RXeMmBsRAbsMc8tdhAoGBAIof
onpuH2HB8vbmVjVT0bJkezxSJL8fBN6KYI4VkscDXWmiZWd1txz5Ieqipo1U9vRS
rRz5LNVJZg7RE9Yg6+aC1r+6LPHJCGHXtcQnOT+/W7+D38/BEACMkijAPlluP+s0
jFxLzblCaC2nusJk/TPqlut+utWptREj64Yqa+WVAoGABasvhISfX3yDRdg/uDrz
GnebVBXeocVZeBFG5g1cLxn9BW5/OteIC35ZucAaWpY0BP+bATJIGLapSHSCZJY+
rtIMhD30LBhwyHyyI4/Z8cB6Qo4uCD/Ra20RGpHwQTnFsHI/NbmlD2xBzAR1Syjl
GHHoQzxwG4RY3hZPIsBu1Rg=
-----END PRIVATE KEY-----
`) // Paste your PEM-encoded private key data here
