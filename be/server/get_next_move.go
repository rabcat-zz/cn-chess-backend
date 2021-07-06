package server

import "context"

func (s *CNChessBackendServer) GetNextMove(
	ctx context.Context,
	board [][]string,
) ([][]string, error) {
	newBoard := [][]string{
		{board[0][1]},
		{board[0][0]},
	}
	return newBoard, nil
}