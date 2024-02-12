package db

func DeletePostFromBD(postId *int, userId *uint64) (int64, error) {
	query := `DELETE FROM posts WHERE post_id = $1 AND user_id = $2`
	result, err := DB.Exec(query, *postId, *userId)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
