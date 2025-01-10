package libs_test

// func xTestCanCreateAValidJWTToken(t *testing.T) {
// 	var masterTokenValue string = domain.UUIDv4()
// 	var duration time.Duration = time.Minute
// 	var issuedAt time.Time = time.Now()
// 	var expireAt time.Time = issuedAt.Add(duration)
// 	log.Println("expireAt", expireAt)

// 	tokener := infraLibs.NewTokenerJWT("any-secret-key-this-is-for-developer-use")
// 	token, err := tokener.Create(masterTokenValue, duration)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, token)
// 	parts := strings.Split(token, ".")
// 	require.Equal(t, 3, len(parts))

// 	validatedPayload, err := tokener.Validate(token)
// 	require.NoError(t, err)
// 	require.NotEmpty(t, validatedPayload)
// 	require.Equal(t, masterTokenValue, validatedPayload.id)
// 	require.WithinDuration(t, issuedAt, validatedPayload.issuedAt, time.Microsecond)
// 	require.WithinDuration(t, expireAt, validatedPayload.expireAt, time.Microsecond)
// }
