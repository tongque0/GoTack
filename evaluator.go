package gotack

import "fmt"

type GameTreeType int

const (
	AlphaBeta GameTreeType = iota // 使用 Alpha-Beta  算法
	PVS                           // 使用 PVS 剪枝算法
	// 可以添加更多的算法类型
)

type Evaluator struct {
	TreeType     GameTreeType
	Depth        int
	IsMaxPlayer  bool
	EvaluateFunc func(board Board, isMaxPlayer bool, opts ...interface{}) float64
}

// NewEvaluator 创建并初始化一个 Evaluator 对象。
// 此函数接收以下参数：
// - treeType: 博弈树的类型。当前只支持 AlphaBeta 类型。
// - depth: 搜索深度，表示算法需要搜索的层次。
// - isMaxPlayer: 表示调用者是否为最大化玩家,先手通常为true。
// - evalFunc: 一个评估函数，用于评估棋盘状态，返回一个表示局面评估值的浮点数。
//   此函数接收棋盘状态和玩家类型，可以接受额外的可选参数。
// 返回值：
// - 返回一个指向 Evaluator 的指针。如果不支持指定的博弈树类型，则返回 nil。
//
// 示例用法：
//   evalFunc := func(board Board, isMaxPlayer bool, opts ...interface{}) float64 {
//       // 实现具体的评估逻辑
//       return 0.0 // 返回评估值
//   }
//   evaluator := NewEvaluator(AlphaBeta, 5, true, evalFunc)
func NewEvaluator(treeType GameTreeType, depth int, isMaxPlayer bool, evalFunc func(board Board, isMaxPlayer bool, opts ...interface{}) float64) *Evaluator {
	if treeType != AlphaBeta {
		fmt.Println("Unsupported tree type")
		return nil
	}
	return &Evaluator{
		TreeType:     treeType,
		Depth:        depth,
		IsMaxPlayer:  isMaxPlayer,
		EvaluateFunc: evalFunc,
	}
}

// GetBestMove 返回最近一次评估中找到的最佳移动。
// 此方法通过在指定的棋盘状态上运行博弈树搜索算法来确定最佳移动。
//
// 参数:
// - board: Board 接口，代表当前棋盘的状态，需要由调用者提供。
//
// 返回值:
// - Move: 从当前棋盘状态中评估得到的最佳移动。
//
// 方法逻辑:
// 根据 Evaluator 结构中的 TreeType 字段，选择适当的博弈树搜索算法。
// - Minimax: 如果树类型为 Minimax，假设 Minimax 方法返回一个 EvaluatedMove 结构，
//            则调用 Minimax 方法并从中提取最佳移动。
// - AlphaBeta: 如果树类型为 AlphaBeta，调用 alphaBeta 方法，并直接返回从该方法获得的最佳移动。
//              alpha 和 beta 的初始值分别设为 -1000000 和 1000000。
// - Other Algorithms: 可以添加其他类型的博弈树算法，每种类型需根据其特定逻辑执行并返回最佳移动。
// - default: 如果树类型不被支持，打印错误消息并返回一个默认的 Move 结构。
//
// 示例用法:
//   board := // 初始化或获取当前棋盘状态
//   evaluator := // 创建并初始化 Evaluator 实例
//   bestMove := evaluator.GetBestMove(board) // 使用 bestMove 进行下一步操作
func (e *Evaluator) GetBestMove(board Board) Move {
	var bestMove Move
	switch e.TreeType {
	case AlphaBeta:
		// 假设 AlphaBeta 方法返回一个 EvaluatedMove 结构
		_, bestMove := e.alphaBeta(board, e.Depth, -1000000, 1000000, e.IsMaxPlayer)
		return bestMove
	case PVS:
		// 假设 AlphaBeta 方法返回一个 EvaluatedMove 结构
		_, bestMove := e.pvs(board, e.Depth, -1000000, 1000000, e.IsMaxPlayer)
		return bestMove
	// 添加其他算法类型的 case 分支
	// 例如：
	// case SomeOtherAlgorithm:
	//     result := e.SomeOtherAlgorithm(board, e.Depth, e.IsMaxPlayer)
	//     bestMove = result.Move

	default:
		fmt.Println("Unsupported tree type")
	}
	return bestMove
}
