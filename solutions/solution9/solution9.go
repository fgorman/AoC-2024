package solution9

import (
	"fmt"
	"slices"
	"strconv"
)

type Block struct {
	Start int
	Size  int
}

const TEST1 = "2333133121414131402"

func Solution(in string) {
	blocks := blocksFromLayout(in)

	blocksShiftBlock := shiftBlocks(blocks)
	checksum := computeChecksum(blocksShiftBlock)
	fmt.Println("Checksum with block shift:", checksum)

	blocksShiftFile := shiftFiles(blocks)
	checksum = computeChecksum(blocksShiftFile)
	fmt.Println("Check with file shift:", checksum)
}

func blocksFromLayout(in string) []int {
	layout := []rune(in)
	blocks := []int{}

	blockId := 0

	for idx := 0; idx < len(layout); idx += 2 {
		numBlocks, _ := strconv.Atoi(string(layout[idx]))
		newBlocks := slices.Repeat([]int{blockId}, numBlocks)
		blocks = append(blocks, newBlocks...)

		if idx >= len(layout)-1 {
			break
		}

		numBlocks, _ = strconv.Atoi(string(layout[idx+1]))
		newBlocks = slices.Repeat([]int{-1}, numBlocks)
		blocks = append(blocks, newBlocks...)

		blockId++
	}

	return blocks
}

func shiftBlocks(blocks []int) []int {
	newBlocks := slices.Clone(blocks)
	l := 0
	r := len(blocks) - 1

	for l < r {
		if newBlocks[l] == -1 && newBlocks[r] != -1 {
			newBlocks[l] = newBlocks[r]
			newBlocks[r] = -1
			l++
			r--
		}

		if newBlocks[l] != -1 {
			l++
		}

		if newBlocks[r] == -1 {
			r--
		}
	}

	return newBlocks
}

func getEmptyLocations(blocks []int) []Block {
	emptyLocs := []Block{}
	newEmpty := Block{}
	inEmptyBlock := false
	for idx, id := range blocks {
		if inEmptyBlock && id != -1 {
			inEmptyBlock = false
			emptyLocs = append(emptyLocs, newEmpty)
			newEmpty = Block{}
		}

		if inEmptyBlock && id == -1 {
			newEmpty.Size++
		}

		if !inEmptyBlock && id == -1 {
			inEmptyBlock = true
			newEmpty.Start = idx
			newEmpty.Size++
		}
	}
	return emptyLocs
}

func getFileBlocks(blocks []int) map[int]Block {
	files := map[int]Block{}

	for idx, id := range blocks {
		if id == -1 {
			continue
		}

		if file, ok := files[id]; ok {
			file.Size++
			files[id] = file
		} else {
			files[id] = Block{idx, 1}
		}
	}

	return files
}

func shiftFiles(blocks []int) []int {
	newBlocks := slices.Clone(blocks)
	emptyBlocks := getEmptyLocations(newBlocks)
	fileBlocks := getFileBlocks(newBlocks)

	for fIdx := len(fileBlocks) - 1; fIdx >= 0; fIdx-- {
		fileBlock := fileBlocks[fIdx]
		for eIdx, empty := range emptyBlocks {
			if empty.Size >= fileBlock.Size && empty.Start < fileBlock.Start {
				newStart := empty.Start
				newEnd := empty.Start + fileBlock.Size
				copy(newBlocks[newStart:newEnd], slices.Repeat([]int{fIdx}, fileBlock.Size))

				originalStart := fileBlock.Start
				originalEnd := fileBlock.Start + fileBlock.Size
				copy(newBlocks[originalStart:originalEnd], slices.Repeat([]int{-1}, fileBlock.Size))

				emptyBlocks[eIdx] = Block{newEnd, empty.Size - fileBlock.Size}

				break
			}
		}
	}

	return newBlocks
}

func computeChecksum(blocks []int) (checksum int) {
	for idx, blockId := range blocks {
		if blockId == -1 {
			continue
		}

		checksum += blockId * idx
	}

	return
}
