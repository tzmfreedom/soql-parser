// Code generated from SOQL.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // SOQL

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 143, 644,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9, 70, 4,
	71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75, 4, 76,
	9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4, 81, 9,
	81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86, 9, 86,
	4, 87, 9, 87, 4, 88, 9, 88, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 181, 10, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 188, 10, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5,
	10, 205, 10, 10, 3, 11, 5, 11, 208, 10, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 221, 10, 12, 3,
	13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 249, 10, 23, 3, 24, 3, 24, 3, 25,
	3, 25, 3, 26, 3, 26, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 5, 28, 262, 10,
	28, 3, 28, 5, 28, 265, 10, 28, 3, 28, 5, 28, 268, 10, 28, 3, 28, 3, 28,
	5, 28, 272, 10, 28, 5, 28, 274, 10, 28, 3, 28, 5, 28, 277, 10, 28, 3, 28,
	5, 28, 280, 10, 28, 3, 28, 5, 28, 283, 10, 28, 3, 28, 5, 28, 286, 10, 28,
	3, 28, 5, 28, 289, 10, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 7,
	29, 297, 10, 29, 12, 29, 14, 29, 300, 11, 29, 3, 30, 3, 30, 3, 30, 3, 30,
	7, 30, 306, 10, 30, 12, 30, 14, 30, 309, 11, 30, 3, 31, 3, 31, 3, 31, 3,
	31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 5, 33, 323,
	10, 33, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 5, 38, 342, 10,
	38, 3, 38, 3, 38, 5, 38, 346, 10, 38, 3, 39, 3, 39, 3, 39, 3, 39, 5, 39,
	352, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 358, 10, 40, 3, 40, 5,
	40, 361, 10, 40, 3, 40, 5, 40, 364, 10, 40, 3, 40, 5, 40, 367, 10, 40,
	3, 40, 5, 40, 370, 10, 40, 3, 40, 5, 40, 373, 10, 40, 3, 40, 5, 40, 376,
	10, 40, 3, 40, 5, 40, 379, 10, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3,
	41, 7, 41, 387, 10, 41, 12, 41, 14, 41, 390, 11, 41, 3, 42, 3, 42, 3, 42,
	3, 42, 5, 42, 396, 10, 42, 3, 43, 3, 43, 5, 43, 400, 10, 43, 3, 44, 3,
	44, 5, 44, 404, 10, 44, 3, 45, 3, 45, 5, 45, 408, 10, 45, 3, 46, 5, 46,
	411, 10, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 5, 47, 418, 10, 47, 3,
	47, 3, 47, 3, 48, 3, 48, 3, 48, 7, 48, 425, 10, 48, 12, 48, 14, 48, 428,
	11, 48, 3, 49, 3, 49, 3, 49, 5, 49, 433, 10, 49, 3, 50, 3, 50, 3, 50, 3,
	50, 5, 50, 439, 10, 50, 3, 50, 3, 50, 3, 51, 6, 51, 444, 10, 51, 13, 51,
	14, 51, 445, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 54, 3,
	54, 3, 54, 3, 55, 3, 55, 3, 55, 7, 55, 461, 10, 55, 12, 55, 14, 55, 464,
	11, 55, 3, 56, 5, 56, 467, 10, 56, 3, 56, 3, 56, 5, 56, 471, 10, 56, 3,
	57, 3, 57, 3, 57, 6, 57, 476, 10, 57, 13, 57, 14, 57, 477, 3, 58, 3, 58,
	3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 5, 58, 490, 10,
	58, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 5, 59, 497, 10, 59, 3, 60, 3, 60,
	3, 60, 7, 60, 502, 10, 60, 12, 60, 14, 60, 505, 11, 60, 3, 61, 5, 61, 508,
	10, 61, 3, 61, 3, 61, 5, 61, 512, 10, 61, 3, 62, 3, 62, 3, 62, 3, 62, 3,
	63, 3, 63, 3, 63, 5, 63, 521, 10, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3, 65,
	3, 65, 3, 65, 3, 65, 5, 65, 531, 10, 65, 3, 66, 3, 66, 3, 66, 3, 66, 3,
	67, 3, 67, 5, 67, 539, 10, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69,
	3, 69, 7, 69, 548, 10, 69, 12, 69, 14, 69, 551, 11, 69, 3, 70, 3, 70, 3,
	70, 5, 70, 556, 10, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 72, 3, 72, 3, 73,
	3, 73, 3, 73, 7, 73, 567, 10, 73, 12, 73, 14, 73, 570, 11, 73, 3, 74, 3,
	74, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 7, 75, 581, 10, 75,
	12, 75, 14, 75, 584, 11, 75, 3, 75, 3, 75, 5, 75, 588, 10, 75, 3, 76, 3,
	76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 78, 3, 78, 3, 78, 3, 79, 3, 79, 3, 79,
	3, 79, 3, 79, 3, 80, 3, 80, 3, 80, 7, 80, 607, 10, 80, 12, 80, 14, 80,
	610, 11, 80, 3, 81, 3, 81, 5, 81, 614, 10, 81, 3, 82, 3, 82, 3, 82, 7,
	82, 619, 10, 82, 12, 82, 14, 82, 622, 11, 82, 3, 83, 3, 83, 5, 83, 626,
	10, 83, 3, 83, 5, 83, 629, 10, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 85, 3,
	86, 3, 86, 5, 86, 638, 10, 86, 3, 87, 3, 87, 3, 88, 3, 88, 3, 88, 2, 2,
	89, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36,
	38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72,
	74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106,
	108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 136,
	138, 140, 142, 144, 146, 148, 150, 152, 154, 156, 158, 160, 162, 164, 166,
	168, 170, 172, 174, 2, 18, 3, 2, 36, 52, 3, 2, 79, 101, 3, 2, 102, 122,
	4, 2, 141, 141, 143, 143, 4, 2, 14, 14, 30, 30, 3, 2, 53, 65, 3, 2, 66,
	71, 3, 2, 72, 73, 3, 2, 74, 78, 4, 2, 6, 6, 27, 27, 3, 2, 4, 5, 3, 2, 36,
	39, 4, 2, 8, 8, 11, 11, 4, 2, 15, 15, 21, 21, 5, 2, 35, 35, 45, 45, 50,
	50, 4, 2, 47, 47, 51, 51, 2, 648, 2, 176, 3, 2, 2, 2, 4, 180, 3, 2, 2,
	2, 6, 187, 3, 2, 2, 2, 8, 189, 3, 2, 2, 2, 10, 191, 3, 2, 2, 2, 12, 193,
	3, 2, 2, 2, 14, 195, 3, 2, 2, 2, 16, 197, 3, 2, 2, 2, 18, 204, 3, 2, 2,
	2, 20, 207, 3, 2, 2, 2, 22, 220, 3, 2, 2, 2, 24, 222, 3, 2, 2, 2, 26, 224,
	3, 2, 2, 2, 28, 226, 3, 2, 2, 2, 30, 230, 3, 2, 2, 2, 32, 232, 3, 2, 2,
	2, 34, 234, 3, 2, 2, 2, 36, 236, 3, 2, 2, 2, 38, 238, 3, 2, 2, 2, 40, 240,
	3, 2, 2, 2, 42, 242, 3, 2, 2, 2, 44, 248, 3, 2, 2, 2, 46, 250, 3, 2, 2,
	2, 48, 252, 3, 2, 2, 2, 50, 254, 3, 2, 2, 2, 52, 256, 3, 2, 2, 2, 54, 258,
	3, 2, 2, 2, 56, 292, 3, 2, 2, 2, 58, 301, 3, 2, 2, 2, 60, 310, 3, 2, 2,
	2, 62, 314, 3, 2, 2, 2, 64, 317, 3, 2, 2, 2, 66, 324, 3, 2, 2, 2, 68, 327,
	3, 2, 2, 2, 70, 331, 3, 2, 2, 2, 72, 334, 3, 2, 2, 2, 74, 337, 3, 2, 2,
	2, 76, 347, 3, 2, 2, 2, 78, 353, 3, 2, 2, 2, 80, 382, 3, 2, 2, 2, 82, 395,
	3, 2, 2, 2, 84, 399, 3, 2, 2, 2, 86, 401, 3, 2, 2, 2, 88, 405, 3, 2, 2,
	2, 90, 410, 3, 2, 2, 2, 92, 414, 3, 2, 2, 2, 94, 421, 3, 2, 2, 2, 96, 432,
	3, 2, 2, 2, 98, 434, 3, 2, 2, 2, 100, 443, 3, 2, 2, 2, 102, 447, 3, 2,
	2, 2, 104, 451, 3, 2, 2, 2, 106, 454, 3, 2, 2, 2, 108, 457, 3, 2, 2, 2,
	110, 466, 3, 2, 2, 2, 112, 475, 3, 2, 2, 2, 114, 489, 3, 2, 2, 2, 116,
	496, 3, 2, 2, 2, 118, 498, 3, 2, 2, 2, 120, 507, 3, 2, 2, 2, 122, 513,
	3, 2, 2, 2, 124, 520, 3, 2, 2, 2, 126, 522, 3, 2, 2, 2, 128, 526, 3, 2,
	2, 2, 130, 532, 3, 2, 2, 2, 132, 538, 3, 2, 2, 2, 134, 540, 3, 2, 2, 2,
	136, 544, 3, 2, 2, 2, 138, 552, 3, 2, 2, 2, 140, 557, 3, 2, 2, 2, 142,
	559, 3, 2, 2, 2, 144, 563, 3, 2, 2, 2, 146, 571, 3, 2, 2, 2, 148, 587,
	3, 2, 2, 2, 150, 589, 3, 2, 2, 2, 152, 591, 3, 2, 2, 2, 154, 593, 3, 2,
	2, 2, 156, 598, 3, 2, 2, 2, 158, 603, 3, 2, 2, 2, 160, 613, 3, 2, 2, 2,
	162, 615, 3, 2, 2, 2, 164, 623, 3, 2, 2, 2, 166, 630, 3, 2, 2, 2, 168,
	632, 3, 2, 2, 2, 170, 637, 3, 2, 2, 2, 172, 639, 3, 2, 2, 2, 174, 641,
	3, 2, 2, 2, 176, 177, 9, 2, 2, 2, 177, 3, 3, 2, 2, 2, 178, 181, 5, 2, 2,
	2, 179, 181, 7, 17, 2, 2, 180, 178, 3, 2, 2, 2, 180, 179, 3, 2, 2, 2, 181,
	5, 3, 2, 2, 2, 182, 188, 7, 138, 2, 2, 183, 188, 5, 4, 3, 2, 184, 188,
	5, 24, 13, 2, 185, 188, 5, 26, 14, 2, 186, 188, 5, 44, 23, 2, 187, 182,
	3, 2, 2, 2, 187, 183, 3, 2, 2, 2, 187, 184, 3, 2, 2, 2, 187, 185, 3, 2,
	2, 2, 187, 186, 3, 2, 2, 2, 188, 7, 3, 2, 2, 2, 189, 190, 5, 6, 4, 2, 190,
	9, 3, 2, 2, 2, 191, 192, 5, 6, 4, 2, 192, 11, 3, 2, 2, 2, 193, 194, 5,
	6, 4, 2, 194, 13, 3, 2, 2, 2, 195, 196, 5, 6, 4, 2, 196, 15, 3, 2, 2, 2,
	197, 198, 5, 6, 4, 2, 198, 17, 3, 2, 2, 2, 199, 205, 7, 138, 2, 2, 200,
	205, 5, 2, 2, 2, 201, 205, 5, 24, 13, 2, 202, 205, 5, 26, 14, 2, 203, 205,
	5, 44, 23, 2, 204, 199, 3, 2, 2, 2, 204, 200, 3, 2, 2, 2, 204, 201, 3,
	2, 2, 2, 204, 202, 3, 2, 2, 2, 204, 203, 3, 2, 2, 2, 205, 19, 3, 2, 2,
	2, 206, 208, 7, 7, 2, 2, 207, 206, 3, 2, 2, 2, 207, 208, 3, 2, 2, 2, 208,
	209, 3, 2, 2, 2, 209, 210, 5, 18, 10, 2, 210, 21, 3, 2, 2, 2, 211, 221,
	5, 24, 13, 2, 212, 221, 5, 28, 15, 2, 213, 221, 5, 30, 16, 2, 214, 221,
	5, 32, 17, 2, 215, 221, 5, 34, 18, 2, 216, 221, 5, 36, 19, 2, 217, 221,
	5, 38, 20, 2, 218, 221, 5, 40, 21, 2, 219, 221, 5, 42, 22, 2, 220, 211,
	3, 2, 2, 2, 220, 212, 3, 2, 2, 2, 220, 213, 3, 2, 2, 2, 220, 214, 3, 2,
	2, 2, 220, 215, 3, 2, 2, 2, 220, 216, 3, 2, 2, 2, 220, 217, 3, 2, 2, 2,
	220, 218, 3, 2, 2, 2, 220, 219, 3, 2, 2, 2, 221, 23, 3, 2, 2, 2, 222, 223,
	9, 3, 2, 2, 223, 25, 3, 2, 2, 2, 224, 225, 9, 4, 2, 2, 225, 27, 3, 2, 2,
	2, 226, 227, 5, 26, 14, 2, 227, 228, 7, 130, 2, 2, 228, 229, 7, 141, 2,
	2, 229, 29, 3, 2, 2, 2, 230, 231, 7, 140, 2, 2, 231, 31, 3, 2, 2, 2, 232,
	233, 7, 139, 2, 2, 233, 33, 3, 2, 2, 2, 234, 235, 9, 5, 2, 2, 235, 35,
	3, 2, 2, 2, 236, 237, 7, 142, 2, 2, 237, 37, 3, 2, 2, 2, 238, 239, 7, 4,
	2, 2, 239, 39, 3, 2, 2, 2, 240, 241, 9, 6, 2, 2, 241, 41, 3, 2, 2, 2, 242,
	243, 7, 25, 2, 2, 243, 43, 3, 2, 2, 2, 244, 249, 5, 46, 24, 2, 245, 249,
	5, 48, 25, 2, 246, 249, 5, 50, 26, 2, 247, 249, 5, 52, 27, 2, 248, 244,
	3, 2, 2, 2, 248, 245, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 248, 247, 3, 2,
	2, 2, 249, 45, 3, 2, 2, 2, 250, 251, 9, 7, 2, 2, 251, 47, 3, 2, 2, 2, 252,
	253, 9, 8, 2, 2, 253, 49, 3, 2, 2, 2, 254, 255, 9, 9, 2, 2, 255, 51, 3,
	2, 2, 2, 256, 257, 9, 10, 2, 2, 257, 53, 3, 2, 2, 2, 258, 259, 5, 56, 29,
	2, 259, 261, 5, 58, 30, 2, 260, 262, 5, 60, 31, 2, 261, 260, 3, 2, 2, 2,
	261, 262, 3, 2, 2, 2, 262, 264, 3, 2, 2, 2, 263, 265, 5, 62, 32, 2, 264,
	263, 3, 2, 2, 2, 264, 265, 3, 2, 2, 2, 265, 267, 3, 2, 2, 2, 266, 268,
	5, 138, 70, 2, 267, 266, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 273, 3,
	2, 2, 2, 269, 271, 5, 64, 33, 2, 270, 272, 5, 66, 34, 2, 271, 270, 3, 2,
	2, 2, 271, 272, 3, 2, 2, 2, 272, 274, 3, 2, 2, 2, 273, 269, 3, 2, 2, 2,
	273, 274, 3, 2, 2, 2, 274, 276, 3, 2, 2, 2, 275, 277, 5, 68, 35, 2, 276,
	275, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 279, 3, 2, 2, 2, 278, 280,
	5, 70, 36, 2, 279, 278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2, 280, 282, 3,
	2, 2, 2, 281, 283, 5, 72, 37, 2, 282, 281, 3, 2, 2, 2, 282, 283, 3, 2,
	2, 2, 283, 285, 3, 2, 2, 2, 284, 286, 5, 74, 38, 2, 285, 284, 3, 2, 2,
	2, 285, 286, 3, 2, 2, 2, 286, 288, 3, 2, 2, 2, 287, 289, 5, 76, 39, 2,
	288, 287, 3, 2, 2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290,
	291, 7, 2, 2, 3, 291, 55, 3, 2, 2, 2, 292, 293, 7, 29, 2, 2, 293, 298,
	5, 82, 42, 2, 294, 295, 7, 132, 2, 2, 295, 297, 5, 82, 42, 2, 296, 294,
	3, 2, 2, 2, 297, 300, 3, 2, 2, 2, 298, 296, 3, 2, 2, 2, 298, 299, 3, 2,
	2, 2, 299, 57, 3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 301, 302, 7, 16, 2, 2,
	302, 307, 5, 110, 56, 2, 303, 304, 7, 132, 2, 2, 304, 306, 5, 110, 56,
	2, 305, 303, 3, 2, 2, 2, 306, 309, 3, 2, 2, 2, 307, 305, 3, 2, 2, 2, 307,
	308, 3, 2, 2, 2, 308, 59, 3, 2, 2, 2, 309, 307, 3, 2, 2, 2, 310, 311, 7,
	31, 2, 2, 311, 312, 7, 46, 2, 2, 312, 313, 5, 12, 7, 2, 313, 61, 3, 2,
	2, 2, 314, 315, 7, 32, 2, 2, 315, 316, 5, 118, 60, 2, 316, 63, 3, 2, 2,
	2, 317, 318, 7, 17, 2, 2, 318, 322, 7, 9, 2, 2, 319, 323, 5, 154, 78, 2,
	320, 323, 5, 156, 79, 2, 321, 323, 5, 152, 77, 2, 322, 319, 3, 2, 2, 2,
	322, 320, 3, 2, 2, 2, 322, 321, 3, 2, 2, 2, 323, 65, 3, 2, 2, 2, 324, 325,
	7, 18, 2, 2, 325, 326, 5, 118, 60, 2, 326, 67, 3, 2, 2, 2, 327, 328, 7,
	44, 2, 2, 328, 329, 7, 9, 2, 2, 329, 330, 5, 162, 82, 2, 330, 69, 3, 2,
	2, 2, 331, 332, 7, 23, 2, 2, 332, 333, 7, 141, 2, 2, 333, 71, 3, 2, 2,
	2, 334, 335, 7, 43, 2, 2, 335, 336, 7, 141, 2, 2, 336, 73, 3, 2, 2, 2,
	337, 338, 7, 34, 2, 2, 338, 341, 5, 172, 87, 2, 339, 340, 7, 132, 2, 2,
	340, 342, 5, 172, 87, 2, 341, 339, 3, 2, 2, 2, 341, 342, 3, 2, 2, 2, 342,
	345, 3, 2, 2, 2, 343, 344, 7, 132, 2, 2, 344, 346, 5, 172, 87, 2, 345,
	343, 3, 2, 2, 2, 345, 346, 3, 2, 2, 2, 346, 75, 3, 2, 2, 2, 347, 348, 7,
	35, 2, 2, 348, 351, 5, 174, 88, 2, 349, 350, 7, 132, 2, 2, 350, 352, 5,
	174, 88, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352, 77, 3, 2,
	2, 2, 353, 354, 7, 135, 2, 2, 354, 355, 5, 80, 41, 2, 355, 357, 5, 58,
	30, 2, 356, 358, 5, 60, 31, 2, 357, 356, 3, 2, 2, 2, 357, 358, 3, 2, 2,
	2, 358, 360, 3, 2, 2, 2, 359, 361, 5, 62, 32, 2, 360, 359, 3, 2, 2, 2,
	360, 361, 3, 2, 2, 2, 361, 363, 3, 2, 2, 2, 362, 364, 5, 138, 70, 2, 363,
	362, 3, 2, 2, 2, 363, 364, 3, 2, 2, 2, 364, 366, 3, 2, 2, 2, 365, 367,
	5, 68, 35, 2, 366, 365, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 369, 3,
	2, 2, 2, 368, 370, 5, 70, 36, 2, 369, 368, 3, 2, 2, 2, 369, 370, 3, 2,
	2, 2, 370, 372, 3, 2, 2, 2, 371, 373, 5, 72, 37, 2, 372, 371, 3, 2, 2,
	2, 372, 373, 3, 2, 2, 2, 373, 375, 3, 2, 2, 2, 374, 376, 5, 74, 38, 2,
	375, 374, 3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 378, 3, 2, 2, 2, 377,
	379, 5, 76, 39, 2, 378, 377, 3, 2, 2, 2, 378, 379, 3, 2, 2, 2, 379, 380,
	3, 2, 2, 2, 380, 381, 7, 134, 2, 2, 381, 79, 3, 2, 2, 2, 382, 383, 7, 29,
	2, 2, 383, 388, 5, 84, 43, 2, 384, 385, 7, 132, 2, 2, 385, 387, 5, 84,
	43, 2, 386, 384, 3, 2, 2, 2, 387, 390, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2,
	388, 389, 3, 2, 2, 2, 389, 81, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 391, 396,
	5, 86, 44, 2, 392, 396, 5, 88, 45, 2, 393, 396, 5, 78, 40, 2, 394, 396,
	5, 98, 50, 2, 395, 391, 3, 2, 2, 2, 395, 392, 3, 2, 2, 2, 395, 393, 3,
	2, 2, 2, 395, 394, 3, 2, 2, 2, 396, 83, 3, 2, 2, 2, 397, 400, 5, 86, 44,
	2, 398, 400, 5, 88, 45, 2, 399, 397, 3, 2, 2, 2, 399, 398, 3, 2, 2, 2,
	400, 85, 3, 2, 2, 2, 401, 403, 5, 90, 46, 2, 402, 404, 5, 20, 11, 2, 403,
	402, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 87, 3, 2, 2, 2, 405, 407, 5,
	92, 47, 2, 406, 408, 5, 20, 11, 2, 407, 406, 3, 2, 2, 2, 407, 408, 3, 2,
	2, 2, 408, 89, 3, 2, 2, 2, 409, 411, 5, 112, 57, 2, 410, 409, 3, 2, 2,
	2, 410, 411, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 5, 10, 6, 2, 413,
	91, 3, 2, 2, 2, 414, 415, 5, 44, 23, 2, 415, 417, 7, 135, 2, 2, 416, 418,
	5, 94, 48, 2, 417, 416, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418, 419, 3,
	2, 2, 2, 419, 420, 7, 134, 2, 2, 420, 93, 3, 2, 2, 2, 421, 426, 5, 96,
	49, 2, 422, 423, 7, 132, 2, 2, 423, 425, 5, 96, 49, 2, 424, 422, 3, 2,
	2, 2, 425, 428, 3, 2, 2, 2, 426, 424, 3, 2, 2, 2, 426, 427, 3, 2, 2, 2,
	427, 95, 3, 2, 2, 2, 428, 426, 3, 2, 2, 2, 429, 433, 5, 90, 46, 2, 430,
	433, 5, 22, 12, 2, 431, 433, 5, 92, 47, 2, 432, 429, 3, 2, 2, 2, 432, 430,
	3, 2, 2, 2, 432, 431, 3, 2, 2, 2, 433, 97, 3, 2, 2, 2, 434, 435, 7, 49,
	2, 2, 435, 436, 5, 90, 46, 2, 436, 438, 5, 100, 51, 2, 437, 439, 5, 106,
	54, 2, 438, 437, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 440, 3, 2, 2, 2,
	440, 441, 7, 42, 2, 2, 441, 99, 3, 2, 2, 2, 442, 444, 5, 102, 52, 2, 443,
	442, 3, 2, 2, 2, 444, 445, 3, 2, 2, 2, 445, 443, 3, 2, 2, 2, 445, 446,
	3, 2, 2, 2, 446, 101, 3, 2, 2, 2, 447, 448, 7, 52, 2, 2, 448, 449, 5, 8,
	5, 2, 449, 450, 5, 104, 53, 2, 450, 103, 3, 2, 2, 2, 451, 452, 7, 48, 2,
	2, 452, 453, 5, 108, 55, 2, 453, 105, 3, 2, 2, 2, 454, 455, 7, 12, 2, 2,
	455, 456, 5, 108, 55, 2, 456, 107, 3, 2, 2, 2, 457, 462, 5, 90, 46, 2,
	458, 459, 7, 132, 2, 2, 459, 461, 5, 90, 46, 2, 460, 458, 3, 2, 2, 2, 461,
	464, 3, 2, 2, 2, 462, 460, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463, 109,
	3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 465, 467, 5, 112, 57, 2, 466, 465, 3,
	2, 2, 2, 466, 467, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 470, 5, 8, 5,
	2, 469, 471, 5, 20, 11, 2, 470, 469, 3, 2, 2, 2, 470, 471, 3, 2, 2, 2,
	471, 111, 3, 2, 2, 2, 472, 473, 5, 8, 5, 2, 473, 474, 7, 129, 2, 2, 474,
	476, 3, 2, 2, 2, 475, 472, 3, 2, 2, 2, 476, 477, 3, 2, 2, 2, 477, 475,
	3, 2, 2, 2, 477, 478, 3, 2, 2, 2, 478, 113, 3, 2, 2, 2, 479, 490, 7, 123,
	2, 2, 480, 490, 7, 124, 2, 2, 481, 490, 7, 125, 2, 2, 482, 490, 7, 126,
	2, 2, 483, 490, 7, 127, 2, 2, 484, 490, 7, 128, 2, 2, 485, 486, 7, 127,
	2, 2, 486, 490, 7, 123, 2, 2, 487, 488, 7, 128, 2, 2, 488, 490, 7, 123,
	2, 2, 489, 479, 3, 2, 2, 2, 489, 480, 3, 2, 2, 2, 489, 481, 3, 2, 2, 2,
	489, 482, 3, 2, 2, 2, 489, 483, 3, 2, 2, 2, 489, 484, 3, 2, 2, 2, 489,
	485, 3, 2, 2, 2, 489, 487, 3, 2, 2, 2, 490, 115, 3, 2, 2, 2, 491, 497,
	7, 19, 2, 2, 492, 493, 7, 24, 2, 2, 493, 497, 7, 19, 2, 2, 494, 497, 7,
	20, 2, 2, 495, 497, 7, 13, 2, 2, 496, 491, 3, 2, 2, 2, 496, 492, 3, 2,
	2, 2, 496, 494, 3, 2, 2, 2, 496, 495, 3, 2, 2, 2, 497, 117, 3, 2, 2, 2,
	498, 503, 5, 120, 61, 2, 499, 500, 9, 11, 2, 2, 500, 502, 5, 120, 61, 2,
	501, 499, 3, 2, 2, 2, 502, 505, 3, 2, 2, 2, 503, 501, 3, 2, 2, 2, 503,
	504, 3, 2, 2, 2, 504, 119, 3, 2, 2, 2, 505, 503, 3, 2, 2, 2, 506, 508,
	7, 24, 2, 2, 507, 506, 3, 2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 511, 3, 2,
	2, 2, 509, 512, 5, 124, 63, 2, 510, 512, 5, 122, 62, 2, 511, 509, 3, 2,
	2, 2, 511, 510, 3, 2, 2, 2, 512, 121, 3, 2, 2, 2, 513, 514, 7, 135, 2,
	2, 514, 515, 5, 118, 60, 2, 515, 516, 7, 134, 2, 2, 516, 123, 3, 2, 2,
	2, 517, 521, 5, 126, 64, 2, 518, 521, 5, 128, 65, 2, 519, 521, 5, 130,
	66, 2, 520, 517, 3, 2, 2, 2, 520, 518, 3, 2, 2, 2, 520, 519, 3, 2, 2, 2,
	521, 125, 3, 2, 2, 2, 522, 523, 5, 132, 67, 2, 523, 524, 5, 114, 58, 2,
	524, 525, 5, 22, 12, 2, 525, 127, 3, 2, 2, 2, 526, 527, 5, 132, 67, 2,
	527, 530, 5, 116, 59, 2, 528, 531, 5, 78, 40, 2, 529, 531, 5, 134, 68,
	2, 530, 528, 3, 2, 2, 2, 530, 529, 3, 2, 2, 2, 531, 129, 3, 2, 2, 2, 532,
	533, 5, 132, 67, 2, 533, 534, 7, 22, 2, 2, 534, 535, 9, 12, 2, 2, 535,
	131, 3, 2, 2, 2, 536, 539, 5, 90, 46, 2, 537, 539, 5, 92, 47, 2, 538, 536,
	3, 2, 2, 2, 538, 537, 3, 2, 2, 2, 539, 133, 3, 2, 2, 2, 540, 541, 7, 135,
	2, 2, 541, 542, 5, 136, 69, 2, 542, 543, 7, 134, 2, 2, 543, 135, 3, 2,
	2, 2, 544, 549, 5, 22, 12, 2, 545, 546, 7, 132, 2, 2, 546, 548, 5, 22,
	12, 2, 547, 545, 3, 2, 2, 2, 548, 551, 3, 2, 2, 2, 549, 547, 3, 2, 2, 2,
	549, 550, 3, 2, 2, 2, 550, 137, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 552,
	555, 7, 33, 2, 2, 553, 556, 5, 140, 71, 2, 554, 556, 5, 142, 72, 2, 555,
	553, 3, 2, 2, 2, 555, 554, 3, 2, 2, 2, 556, 139, 3, 2, 2, 2, 557, 558,
	5, 126, 64, 2, 558, 141, 3, 2, 2, 2, 559, 560, 7, 41, 2, 2, 560, 561, 7,
	40, 2, 2, 561, 562, 5, 144, 73, 2, 562, 143, 3, 2, 2, 2, 563, 568, 5, 146,
	74, 2, 564, 565, 7, 6, 2, 2, 565, 567, 5, 146, 74, 2, 566, 564, 3, 2, 2,
	2, 567, 570, 3, 2, 2, 2, 568, 566, 3, 2, 2, 2, 568, 569, 3, 2, 2, 2, 569,
	145, 3, 2, 2, 2, 570, 568, 3, 2, 2, 2, 571, 572, 5, 14, 8, 2, 572, 573,
	5, 150, 76, 2, 573, 574, 5, 148, 75, 2, 574, 147, 3, 2, 2, 2, 575, 588,
	5, 16, 9, 2, 576, 577, 7, 135, 2, 2, 577, 582, 5, 16, 9, 2, 578, 579, 7,
	132, 2, 2, 579, 581, 5, 16, 9, 2, 580, 578, 3, 2, 2, 2, 581, 584, 3, 2,
	2, 2, 582, 580, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 585, 3, 2, 2, 2,
	584, 582, 3, 2, 2, 2, 585, 586, 7, 134, 2, 2, 586, 588, 3, 2, 2, 2, 587,
	575, 3, 2, 2, 2, 587, 576, 3, 2, 2, 2, 588, 149, 3, 2, 2, 2, 589, 590,
	9, 13, 2, 2, 590, 151, 3, 2, 2, 2, 591, 592, 5, 158, 80, 2, 592, 153, 3,
	2, 2, 2, 593, 594, 7, 28, 2, 2, 594, 595, 7, 135, 2, 2, 595, 596, 5, 158,
	80, 2, 596, 597, 7, 134, 2, 2, 597, 155, 3, 2, 2, 2, 598, 599, 7, 10, 2,
	2, 599, 600, 7, 135, 2, 2, 600, 601, 5, 158, 80, 2, 601, 602, 7, 134, 2,
	2, 602, 157, 3, 2, 2, 2, 603, 608, 5, 160, 81, 2, 604, 605, 7, 132, 2,
	2, 605, 607, 5, 160, 81, 2, 606, 604, 3, 2, 2, 2, 607, 610, 3, 2, 2, 2,
	608, 606, 3, 2, 2, 2, 608, 609, 3, 2, 2, 2, 609, 159, 3, 2, 2, 2, 610,
	608, 3, 2, 2, 2, 611, 614, 5, 90, 46, 2, 612, 614, 5, 92, 47, 2, 613, 611,
	3, 2, 2, 2, 613, 612, 3, 2, 2, 2, 614, 161, 3, 2, 2, 2, 615, 620, 5, 164,
	83, 2, 616, 617, 7, 132, 2, 2, 617, 619, 5, 164, 83, 2, 618, 616, 3, 2,
	2, 2, 619, 622, 3, 2, 2, 2, 620, 618, 3, 2, 2, 2, 620, 621, 3, 2, 2, 2,
	621, 163, 3, 2, 2, 2, 622, 620, 3, 2, 2, 2, 623, 625, 5, 170, 86, 2, 624,
	626, 5, 166, 84, 2, 625, 624, 3, 2, 2, 2, 625, 626, 3, 2, 2, 2, 626, 628,
	3, 2, 2, 2, 627, 629, 5, 168, 85, 2, 628, 627, 3, 2, 2, 2, 628, 629, 3,
	2, 2, 2, 629, 165, 3, 2, 2, 2, 630, 631, 9, 14, 2, 2, 631, 167, 3, 2, 2,
	2, 632, 633, 7, 26, 2, 2, 633, 634, 9, 15, 2, 2, 634, 169, 3, 2, 2, 2,
	635, 638, 5, 90, 46, 2, 636, 638, 5, 92, 47, 2, 637, 635, 3, 2, 2, 2, 637,
	636, 3, 2, 2, 2, 638, 171, 3, 2, 2, 2, 639, 640, 9, 16, 2, 2, 640, 173,
	3, 2, 2, 2, 641, 642, 9, 17, 2, 2, 642, 175, 3, 2, 2, 2, 66, 180, 187,
	204, 207, 220, 248, 261, 264, 267, 271, 273, 276, 279, 282, 285, 288, 298,
	307, 322, 341, 345, 351, 357, 360, 363, 366, 369, 372, 375, 378, 388, 395,
	399, 403, 407, 410, 417, 426, 432, 438, 445, 462, 466, 470, 477, 489, 496,
	503, 507, 511, 520, 530, 538, 549, 555, 568, 582, 587, 608, 613, 620, 625,
	628, 637,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "'='", "", "'<='",
	"'>='", "'>'", "'<'", "'.'", "':'", "';'", "','", "'*'", "')'", "'('",
	"'+'", "'-'",
}
var symbolicNames = []string{
	"", "WHITE_SPACE", "STRING_VALUE", "LIKE_STRING_VALUE", "AND", "AS", "ASC",
	"BY", "CUBE", "DESC", "ELSE", "EXCLUDES", "FALSE", "FIRST", "FROM", "GROUP",
	"HAVING", "IN", "INCLUDES", "LAST", "LIKE", "LIMIT", "NOT", "NULL", "NULLS",
	"OR", "ROLLUP", "SELECT", "TRUE", "USING", "WHERE", "WITH", "FOR", "UPDATE",
	"ABOVE", "ABOVE_OR_BELOW", "AT", "BELOW", "CATEGORY", "DATA", "END", "OFFSET",
	"ORDER", "REFERENCE", "SCOPE", "TRACKING", "THEN", "TYPEOF", "VIEW", "VIEWSTAT",
	"WHEN", "CALENDAR_MONTH", "CALENDAR_QUARTER", "CALENDAR_YEAR", "DAY_IN_MONTH",
	"DAY_IN_WEEK", "DAY_IN_YEAR", "DAY_ONLY", "FISCAL_MONTH", "FISCAL_QUARTER",
	"FISCAL_YEAR", "HOUR_IN_DAY", "WEEK_IN_MONTH", "WEEK_IN_YEAR", "AVG", "COUNT",
	"COUNT_DISTINCT", "MIN", "MAX", "SUM", "DISTANCE", "GEOLOCATION", "FORMAT",
	"TOLABEL", "CONVERT_TIME_ZONE", "CONVERT_CURRENCY", "GROUPING", "YESTERDAY",
	"TODAY", "TOMORROW", "LAST_WEEK", "THIS_WEEK", "NEXT_WEEK", "LAST_MONTH",
	"THIS_MONTH", "NEXT_MONTH", "LAST_90_DAYS", "NEXT_90_DAYS", "THIS_QUARTER",
	"LAST_QUARTER", "NEXT_QUARTER", "THIS_YEAR", "LAST_YEAR", "NEXT_YEAR",
	"THIS_FISCAL_QUARTER", "LAST_FISCAL_QUARTER", "NEXT_FISCAL_QUARTER", "THIS_FISCAL_YEAR",
	"LAST_FISCAL_YEAR", "NEXT_FISCAL_YEAR", "NEXT_N_DAYS", "LAST_N_DAYS", "N_DAYS_AGO",
	"NEXT_N_WEEKS", "LAST_N_WEEKS", "N_WEEKS_AGO", "NEXT_N_MONTHS", "LAST_N_MONTHS",
	"N_MONTHS_AGO", "NEXT_N_QUARTERS", "LAST_N_QUARTERS", "N_QUARTERS_AGO",
	"NEXT_N_YEARS", "LAST_N_YEARS", "N_YEARS_AGO", "NEXT_N_FISCAL_QUARTERS",
	"LAST_N_FISCAL_QUARTERS", "N_FISCAL_QUARTERS_AGO", "NEXT_N_FISCAL_YEARS",
	"LAST_N_FISCAL_YEARS", "N_FISCAL_YEARS_AGO", "EQ", "NOT_EQ", "LET", "GET",
	"GTH", "LTH", "DOT", "COLON", "SEMICOLON", "COMMA", "ASTERISK", "RPAREN",
	"LPAREN", "PLUS", "MINUS", "ID", "DATE", "DATETIME", "UNSIGNED_INTEGER",
	"REAL_NUMBER", "SIGNED_INTEGER",
}

var ruleNames = []string{
	"keywords_alias_allowed", "keywords_name_allowed", "name", "object_name",
	"field_name", "filter_scope_name", "data_category_group_name", "data_category_name",
	"alias_name", "alias", "literal", "date_formula_literal", "date_formula_n_literal_name",
	"date_formula_n_literal", "datetime_literal", "date_literal", "integer_literal",
	"real_literal", "string_literal", "boolean_literal", "null_literal", "function_name",
	"function_date", "function_aggregate", "function_location", "function_other",
	"soql_query", "select_clause", "from_clause", "using_clause", "where_clause",
	"groupby_clause", "having_clause", "orderby_clause", "limit_clause", "offset_clause",
	"for_clause", "update_clause", "soql_subquery", "subquery_select_clause",
	"select_spec", "subquery_select_spec", "field_spec", "function_call_spec",
	"field", "function_call", "function_parameter_list", "function_parameter",
	"typeof_spec", "typeof_when_then_clause_list", "typeof_when_then_clause",
	"typeof_then_clause", "typeof_else_clause", "field_list", "object_spec",
	"object_prefix", "comparison_operator", "set_operator", "condition", "condition1",
	"parenthesis", "simple_condition", "field_based_condition", "set_based_condition",
	"like_based_condition", "condition_field", "set_values", "set_value_list",
	"with_clause", "with_plain_clause", "with_data_category_clause", "data_category_spec_list",
	"data_category_spec", "data_category_parameter_list", "data_category_selector",
	"group_by_plain_clause", "group_by_rollup_clause", "group_by_cube_clause",
	"group_by_list", "group_by_spec", "order_by_list", "order_by_spec", "order_by_direction_clause",
	"order_by_nulls_clause", "order_by_field", "for_value", "update_value",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SOQLParser struct {
	*antlr.BaseParser
}

func NewSOQLParser(input antlr.TokenStream) *SOQLParser {
	this := new(SOQLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SOQL.g4"

	return this
}

// SOQLParser tokens.
const (
	SOQLParserEOF                    = antlr.TokenEOF
	SOQLParserWHITE_SPACE            = 1
	SOQLParserSTRING_VALUE           = 2
	SOQLParserLIKE_STRING_VALUE      = 3
	SOQLParserAND                    = 4
	SOQLParserAS                     = 5
	SOQLParserASC                    = 6
	SOQLParserBY                     = 7
	SOQLParserCUBE                   = 8
	SOQLParserDESC                   = 9
	SOQLParserELSE                   = 10
	SOQLParserEXCLUDES               = 11
	SOQLParserFALSE                  = 12
	SOQLParserFIRST                  = 13
	SOQLParserFROM                   = 14
	SOQLParserGROUP                  = 15
	SOQLParserHAVING                 = 16
	SOQLParserIN                     = 17
	SOQLParserINCLUDES               = 18
	SOQLParserLAST                   = 19
	SOQLParserLIKE                   = 20
	SOQLParserLIMIT                  = 21
	SOQLParserNOT                    = 22
	SOQLParserNULL                   = 23
	SOQLParserNULLS                  = 24
	SOQLParserOR                     = 25
	SOQLParserROLLUP                 = 26
	SOQLParserSELECT                 = 27
	SOQLParserTRUE                   = 28
	SOQLParserUSING                  = 29
	SOQLParserWHERE                  = 30
	SOQLParserWITH                   = 31
	SOQLParserFOR                    = 32
	SOQLParserUPDATE                 = 33
	SOQLParserABOVE                  = 34
	SOQLParserABOVE_OR_BELOW         = 35
	SOQLParserAT                     = 36
	SOQLParserBELOW                  = 37
	SOQLParserCATEGORY               = 38
	SOQLParserDATA                   = 39
	SOQLParserEND                    = 40
	SOQLParserOFFSET                 = 41
	SOQLParserORDER                  = 42
	SOQLParserREFERENCE              = 43
	SOQLParserSCOPE                  = 44
	SOQLParserTRACKING               = 45
	SOQLParserTHEN                   = 46
	SOQLParserTYPEOF                 = 47
	SOQLParserVIEW                   = 48
	SOQLParserVIEWSTAT               = 49
	SOQLParserWHEN                   = 50
	SOQLParserCALENDAR_MONTH         = 51
	SOQLParserCALENDAR_QUARTER       = 52
	SOQLParserCALENDAR_YEAR          = 53
	SOQLParserDAY_IN_MONTH           = 54
	SOQLParserDAY_IN_WEEK            = 55
	SOQLParserDAY_IN_YEAR            = 56
	SOQLParserDAY_ONLY               = 57
	SOQLParserFISCAL_MONTH           = 58
	SOQLParserFISCAL_QUARTER         = 59
	SOQLParserFISCAL_YEAR            = 60
	SOQLParserHOUR_IN_DAY            = 61
	SOQLParserWEEK_IN_MONTH          = 62
	SOQLParserWEEK_IN_YEAR           = 63
	SOQLParserAVG                    = 64
	SOQLParserCOUNT                  = 65
	SOQLParserCOUNT_DISTINCT         = 66
	SOQLParserMIN                    = 67
	SOQLParserMAX                    = 68
	SOQLParserSUM                    = 69
	SOQLParserDISTANCE               = 70
	SOQLParserGEOLOCATION            = 71
	SOQLParserFORMAT                 = 72
	SOQLParserTOLABEL                = 73
	SOQLParserCONVERT_TIME_ZONE      = 74
	SOQLParserCONVERT_CURRENCY       = 75
	SOQLParserGROUPING               = 76
	SOQLParserYESTERDAY              = 77
	SOQLParserTODAY                  = 78
	SOQLParserTOMORROW               = 79
	SOQLParserLAST_WEEK              = 80
	SOQLParserTHIS_WEEK              = 81
	SOQLParserNEXT_WEEK              = 82
	SOQLParserLAST_MONTH             = 83
	SOQLParserTHIS_MONTH             = 84
	SOQLParserNEXT_MONTH             = 85
	SOQLParserLAST_90_DAYS           = 86
	SOQLParserNEXT_90_DAYS           = 87
	SOQLParserTHIS_QUARTER           = 88
	SOQLParserLAST_QUARTER           = 89
	SOQLParserNEXT_QUARTER           = 90
	SOQLParserTHIS_YEAR              = 91
	SOQLParserLAST_YEAR              = 92
	SOQLParserNEXT_YEAR              = 93
	SOQLParserTHIS_FISCAL_QUARTER    = 94
	SOQLParserLAST_FISCAL_QUARTER    = 95
	SOQLParserNEXT_FISCAL_QUARTER    = 96
	SOQLParserTHIS_FISCAL_YEAR       = 97
	SOQLParserLAST_FISCAL_YEAR       = 98
	SOQLParserNEXT_FISCAL_YEAR       = 99
	SOQLParserNEXT_N_DAYS            = 100
	SOQLParserLAST_N_DAYS            = 101
	SOQLParserN_DAYS_AGO             = 102
	SOQLParserNEXT_N_WEEKS           = 103
	SOQLParserLAST_N_WEEKS           = 104
	SOQLParserN_WEEKS_AGO            = 105
	SOQLParserNEXT_N_MONTHS          = 106
	SOQLParserLAST_N_MONTHS          = 107
	SOQLParserN_MONTHS_AGO           = 108
	SOQLParserNEXT_N_QUARTERS        = 109
	SOQLParserLAST_N_QUARTERS        = 110
	SOQLParserN_QUARTERS_AGO         = 111
	SOQLParserNEXT_N_YEARS           = 112
	SOQLParserLAST_N_YEARS           = 113
	SOQLParserN_YEARS_AGO            = 114
	SOQLParserNEXT_N_FISCAL_QUARTERS = 115
	SOQLParserLAST_N_FISCAL_QUARTERS = 116
	SOQLParserN_FISCAL_QUARTERS_AGO  = 117
	SOQLParserNEXT_N_FISCAL_YEARS    = 118
	SOQLParserLAST_N_FISCAL_YEARS    = 119
	SOQLParserN_FISCAL_YEARS_AGO     = 120
	SOQLParserEQ                     = 121
	SOQLParserNOT_EQ                 = 122
	SOQLParserLET                    = 123
	SOQLParserGET                    = 124
	SOQLParserGTH                    = 125
	SOQLParserLTH                    = 126
	SOQLParserDOT                    = 127
	SOQLParserCOLON                  = 128
	SOQLParserSEMICOLON              = 129
	SOQLParserCOMMA                  = 130
	SOQLParserASTERISK               = 131
	SOQLParserRPAREN                 = 132
	SOQLParserLPAREN                 = 133
	SOQLParserPLUS                   = 134
	SOQLParserMINUS                  = 135
	SOQLParserID                     = 136
	SOQLParserDATE                   = 137
	SOQLParserDATETIME               = 138
	SOQLParserUNSIGNED_INTEGER       = 139
	SOQLParserREAL_NUMBER            = 140
	SOQLParserSIGNED_INTEGER         = 141
)

// SOQLParser rules.
const (
	SOQLParserRULE_keywords_alias_allowed       = 0
	SOQLParserRULE_keywords_name_allowed        = 1
	SOQLParserRULE_name                         = 2
	SOQLParserRULE_object_name                  = 3
	SOQLParserRULE_field_name                   = 4
	SOQLParserRULE_filter_scope_name            = 5
	SOQLParserRULE_data_category_group_name     = 6
	SOQLParserRULE_data_category_name           = 7
	SOQLParserRULE_alias_name                   = 8
	SOQLParserRULE_alias                        = 9
	SOQLParserRULE_literal                      = 10
	SOQLParserRULE_date_formula_literal         = 11
	SOQLParserRULE_date_formula_n_literal_name  = 12
	SOQLParserRULE_date_formula_n_literal       = 13
	SOQLParserRULE_datetime_literal             = 14
	SOQLParserRULE_date_literal                 = 15
	SOQLParserRULE_integer_literal              = 16
	SOQLParserRULE_real_literal                 = 17
	SOQLParserRULE_string_literal               = 18
	SOQLParserRULE_boolean_literal              = 19
	SOQLParserRULE_null_literal                 = 20
	SOQLParserRULE_function_name                = 21
	SOQLParserRULE_function_date                = 22
	SOQLParserRULE_function_aggregate           = 23
	SOQLParserRULE_function_location            = 24
	SOQLParserRULE_function_other               = 25
	SOQLParserRULE_soql_query                   = 26
	SOQLParserRULE_select_clause                = 27
	SOQLParserRULE_from_clause                  = 28
	SOQLParserRULE_using_clause                 = 29
	SOQLParserRULE_where_clause                 = 30
	SOQLParserRULE_groupby_clause               = 31
	SOQLParserRULE_having_clause                = 32
	SOQLParserRULE_orderby_clause               = 33
	SOQLParserRULE_limit_clause                 = 34
	SOQLParserRULE_offset_clause                = 35
	SOQLParserRULE_for_clause                   = 36
	SOQLParserRULE_update_clause                = 37
	SOQLParserRULE_soql_subquery                = 38
	SOQLParserRULE_subquery_select_clause       = 39
	SOQLParserRULE_select_spec                  = 40
	SOQLParserRULE_subquery_select_spec         = 41
	SOQLParserRULE_field_spec                   = 42
	SOQLParserRULE_function_call_spec           = 43
	SOQLParserRULE_field                        = 44
	SOQLParserRULE_function_call                = 45
	SOQLParserRULE_function_parameter_list      = 46
	SOQLParserRULE_function_parameter           = 47
	SOQLParserRULE_typeof_spec                  = 48
	SOQLParserRULE_typeof_when_then_clause_list = 49
	SOQLParserRULE_typeof_when_then_clause      = 50
	SOQLParserRULE_typeof_then_clause           = 51
	SOQLParserRULE_typeof_else_clause           = 52
	SOQLParserRULE_field_list                   = 53
	SOQLParserRULE_object_spec                  = 54
	SOQLParserRULE_object_prefix                = 55
	SOQLParserRULE_comparison_operator          = 56
	SOQLParserRULE_set_operator                 = 57
	SOQLParserRULE_condition                    = 58
	SOQLParserRULE_condition1                   = 59
	SOQLParserRULE_parenthesis                  = 60
	SOQLParserRULE_simple_condition             = 61
	SOQLParserRULE_field_based_condition        = 62
	SOQLParserRULE_set_based_condition          = 63
	SOQLParserRULE_like_based_condition         = 64
	SOQLParserRULE_condition_field              = 65
	SOQLParserRULE_set_values                   = 66
	SOQLParserRULE_set_value_list               = 67
	SOQLParserRULE_with_clause                  = 68
	SOQLParserRULE_with_plain_clause            = 69
	SOQLParserRULE_with_data_category_clause    = 70
	SOQLParserRULE_data_category_spec_list      = 71
	SOQLParserRULE_data_category_spec           = 72
	SOQLParserRULE_data_category_parameter_list = 73
	SOQLParserRULE_data_category_selector       = 74
	SOQLParserRULE_group_by_plain_clause        = 75
	SOQLParserRULE_group_by_rollup_clause       = 76
	SOQLParserRULE_group_by_cube_clause         = 77
	SOQLParserRULE_group_by_list                = 78
	SOQLParserRULE_group_by_spec                = 79
	SOQLParserRULE_order_by_list                = 80
	SOQLParserRULE_order_by_spec                = 81
	SOQLParserRULE_order_by_direction_clause    = 82
	SOQLParserRULE_order_by_nulls_clause        = 83
	SOQLParserRULE_order_by_field               = 84
	SOQLParserRULE_for_value                    = 85
	SOQLParserRULE_update_value                 = 86
)

// IKeywords_alias_allowedContext is an interface to support dynamic dispatch.
type IKeywords_alias_allowedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywords_alias_allowedContext differentiates from other interfaces.
	IsKeywords_alias_allowedContext()
}

type Keywords_alias_allowedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywords_alias_allowedContext() *Keywords_alias_allowedContext {
	var p = new(Keywords_alias_allowedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_keywords_alias_allowed
	return p
}

func (*Keywords_alias_allowedContext) IsKeywords_alias_allowedContext() {}

func NewKeywords_alias_allowedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Keywords_alias_allowedContext {
	var p = new(Keywords_alias_allowedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_keywords_alias_allowed

	return p
}

func (s *Keywords_alias_allowedContext) GetParser() antlr.Parser { return s.parser }

func (s *Keywords_alias_allowedContext) ABOVE() antlr.TerminalNode {
	return s.GetToken(SOQLParserABOVE, 0)
}

func (s *Keywords_alias_allowedContext) ABOVE_OR_BELOW() antlr.TerminalNode {
	return s.GetToken(SOQLParserABOVE_OR_BELOW, 0)
}

func (s *Keywords_alias_allowedContext) AT() antlr.TerminalNode {
	return s.GetToken(SOQLParserAT, 0)
}

func (s *Keywords_alias_allowedContext) BELOW() antlr.TerminalNode {
	return s.GetToken(SOQLParserBELOW, 0)
}

func (s *Keywords_alias_allowedContext) CATEGORY() antlr.TerminalNode {
	return s.GetToken(SOQLParserCATEGORY, 0)
}

func (s *Keywords_alias_allowedContext) DATA() antlr.TerminalNode {
	return s.GetToken(SOQLParserDATA, 0)
}

func (s *Keywords_alias_allowedContext) END() antlr.TerminalNode {
	return s.GetToken(SOQLParserEND, 0)
}

func (s *Keywords_alias_allowedContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SOQLParserOFFSET, 0)
}

func (s *Keywords_alias_allowedContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SOQLParserORDER, 0)
}

func (s *Keywords_alias_allowedContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(SOQLParserREFERENCE, 0)
}

func (s *Keywords_alias_allowedContext) SCOPE() antlr.TerminalNode {
	return s.GetToken(SOQLParserSCOPE, 0)
}

func (s *Keywords_alias_allowedContext) TRACKING() antlr.TerminalNode {
	return s.GetToken(SOQLParserTRACKING, 0)
}

func (s *Keywords_alias_allowedContext) THEN() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHEN, 0)
}

func (s *Keywords_alias_allowedContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(SOQLParserTYPEOF, 0)
}

func (s *Keywords_alias_allowedContext) VIEW() antlr.TerminalNode {
	return s.GetToken(SOQLParserVIEW, 0)
}

func (s *Keywords_alias_allowedContext) VIEWSTAT() antlr.TerminalNode {
	return s.GetToken(SOQLParserVIEWSTAT, 0)
}

func (s *Keywords_alias_allowedContext) WHEN() antlr.TerminalNode {
	return s.GetToken(SOQLParserWHEN, 0)
}

func (s *Keywords_alias_allowedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Keywords_alias_allowedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Keywords_alias_allowedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterKeywords_alias_allowed(s)
	}
}

func (s *Keywords_alias_allowedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitKeywords_alias_allowed(s)
	}
}

func (s *Keywords_alias_allowedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitKeywords_alias_allowed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Keywords_alias_allowed() (localctx IKeywords_alias_allowedContext) {
	localctx = NewKeywords_alias_allowedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SOQLParserRULE_keywords_alias_allowed)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SOQLParserABOVE-34))|(1<<(SOQLParserABOVE_OR_BELOW-34))|(1<<(SOQLParserAT-34))|(1<<(SOQLParserBELOW-34))|(1<<(SOQLParserCATEGORY-34))|(1<<(SOQLParserDATA-34))|(1<<(SOQLParserEND-34))|(1<<(SOQLParserOFFSET-34))|(1<<(SOQLParserORDER-34))|(1<<(SOQLParserREFERENCE-34))|(1<<(SOQLParserSCOPE-34))|(1<<(SOQLParserTRACKING-34))|(1<<(SOQLParserTHEN-34))|(1<<(SOQLParserTYPEOF-34))|(1<<(SOQLParserVIEW-34))|(1<<(SOQLParserVIEWSTAT-34))|(1<<(SOQLParserWHEN-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IKeywords_name_allowedContext is an interface to support dynamic dispatch.
type IKeywords_name_allowedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywords_name_allowedContext differentiates from other interfaces.
	IsKeywords_name_allowedContext()
}

type Keywords_name_allowedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywords_name_allowedContext() *Keywords_name_allowedContext {
	var p = new(Keywords_name_allowedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_keywords_name_allowed
	return p
}

func (*Keywords_name_allowedContext) IsKeywords_name_allowedContext() {}

func NewKeywords_name_allowedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Keywords_name_allowedContext {
	var p = new(Keywords_name_allowedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_keywords_name_allowed

	return p
}

func (s *Keywords_name_allowedContext) GetParser() antlr.Parser { return s.parser }

func (s *Keywords_name_allowedContext) Keywords_alias_allowed() IKeywords_alias_allowedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywords_alias_allowedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywords_alias_allowedContext)
}

func (s *Keywords_name_allowedContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SOQLParserGROUP, 0)
}

func (s *Keywords_name_allowedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Keywords_name_allowedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Keywords_name_allowedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterKeywords_name_allowed(s)
	}
}

func (s *Keywords_name_allowedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitKeywords_name_allowed(s)
	}
}

func (s *Keywords_name_allowedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitKeywords_name_allowed(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Keywords_name_allowed() (localctx IKeywords_name_allowedContext) {
	localctx = NewKeywords_name_allowedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SOQLParserRULE_keywords_name_allowed)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)
			p.Keywords_alias_allowed()
		}

	case SOQLParserGROUP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(177)
			p.Match(SOQLParserGROUP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) ID() antlr.TerminalNode {
	return s.GetToken(SOQLParserID, 0)
}

func (s *NameContext) Keywords_name_allowed() IKeywords_name_allowedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywords_name_allowedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywords_name_allowedContext)
}

func (s *NameContext) Date_formula_literal() IDate_formula_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_literalContext)
}

func (s *NameContext) Date_formula_n_literal_name() IDate_formula_n_literal_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_n_literal_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_n_literal_nameContext)
}

func (s *NameContext) Function_name() IFunction_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SOQLParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(185)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(SOQLParserID)
		}

	case SOQLParserGROUP, SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)
			p.Keywords_name_allowed()
		}

	case SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(182)
			p.Date_formula_literal()
		}

	case SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Date_formula_n_literal_name()
		}

	case SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR, SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM, SOQLParserDISTANCE, SOQLParserGEOLOCATION, SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(184)
			p.Function_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IObject_nameContext is an interface to support dynamic dispatch.
type IObject_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_nameContext differentiates from other interfaces.
	IsObject_nameContext()
}

type Object_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_nameContext() *Object_nameContext {
	var p = new(Object_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_object_name
	return p
}

func (*Object_nameContext) IsObject_nameContext() {}

func NewObject_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_nameContext {
	var p = new(Object_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_object_name

	return p
}

func (s *Object_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_nameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Object_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterObject_name(s)
	}
}

func (s *Object_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitObject_name(s)
	}
}

func (s *Object_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitObject_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Object_name() (localctx IObject_nameContext) {
	localctx = NewObject_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SOQLParserRULE_object_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(187)
		p.Name()
	}

	return localctx
}

// IField_nameContext is an interface to support dynamic dispatch.
type IField_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_nameContext differentiates from other interfaces.
	IsField_nameContext()
}

type Field_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_nameContext() *Field_nameContext {
	var p = new(Field_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_field_name
	return p
}

func (*Field_nameContext) IsField_nameContext() {}

func NewField_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_nameContext {
	var p = new(Field_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_field_name

	return p
}

func (s *Field_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_nameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Field_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterField_name(s)
	}
}

func (s *Field_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitField_name(s)
	}
}

func (s *Field_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitField_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Field_name() (localctx IField_nameContext) {
	localctx = NewField_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SOQLParserRULE_field_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Name()
	}

	return localctx
}

// IFilter_scope_nameContext is an interface to support dynamic dispatch.
type IFilter_scope_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFilter_scope_nameContext differentiates from other interfaces.
	IsFilter_scope_nameContext()
}

type Filter_scope_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_scope_nameContext() *Filter_scope_nameContext {
	var p = new(Filter_scope_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_filter_scope_name
	return p
}

func (*Filter_scope_nameContext) IsFilter_scope_nameContext() {}

func NewFilter_scope_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_scope_nameContext {
	var p = new(Filter_scope_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_filter_scope_name

	return p
}

func (s *Filter_scope_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Filter_scope_nameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Filter_scope_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_scope_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_scope_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFilter_scope_name(s)
	}
}

func (s *Filter_scope_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFilter_scope_name(s)
	}
}

func (s *Filter_scope_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFilter_scope_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Filter_scope_name() (localctx IFilter_scope_nameContext) {
	localctx = NewFilter_scope_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SOQLParserRULE_filter_scope_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(191)
		p.Name()
	}

	return localctx
}

// IData_category_group_nameContext is an interface to support dynamic dispatch.
type IData_category_group_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_group_nameContext differentiates from other interfaces.
	IsData_category_group_nameContext()
}

type Data_category_group_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_group_nameContext() *Data_category_group_nameContext {
	var p = new(Data_category_group_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_group_name
	return p
}

func (*Data_category_group_nameContext) IsData_category_group_nameContext() {}

func NewData_category_group_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_group_nameContext {
	var p = new(Data_category_group_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_group_name

	return p
}

func (s *Data_category_group_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_group_nameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Data_category_group_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_group_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_group_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_group_name(s)
	}
}

func (s *Data_category_group_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_group_name(s)
	}
}

func (s *Data_category_group_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_group_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_group_name() (localctx IData_category_group_nameContext) {
	localctx = NewData_category_group_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SOQLParserRULE_data_category_group_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Name()
	}

	return localctx
}

// IData_category_nameContext is an interface to support dynamic dispatch.
type IData_category_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_nameContext differentiates from other interfaces.
	IsData_category_nameContext()
}

type Data_category_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_nameContext() *Data_category_nameContext {
	var p = new(Data_category_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_name
	return p
}

func (*Data_category_nameContext) IsData_category_nameContext() {}

func NewData_category_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_nameContext {
	var p = new(Data_category_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_name

	return p
}

func (s *Data_category_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_nameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Data_category_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_name(s)
	}
}

func (s *Data_category_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_name(s)
	}
}

func (s *Data_category_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_name() (localctx IData_category_nameContext) {
	localctx = NewData_category_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SOQLParserRULE_data_category_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(195)
		p.Name()
	}

	return localctx
}

// IAlias_nameContext is an interface to support dynamic dispatch.
type IAlias_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAlias_nameContext differentiates from other interfaces.
	IsAlias_nameContext()
}

type Alias_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAlias_nameContext() *Alias_nameContext {
	var p = new(Alias_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_alias_name
	return p
}

func (*Alias_nameContext) IsAlias_nameContext() {}

func NewAlias_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Alias_nameContext {
	var p = new(Alias_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_alias_name

	return p
}

func (s *Alias_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Alias_nameContext) ID() antlr.TerminalNode {
	return s.GetToken(SOQLParserID, 0)
}

func (s *Alias_nameContext) Keywords_alias_allowed() IKeywords_alias_allowedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IKeywords_alias_allowedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IKeywords_alias_allowedContext)
}

func (s *Alias_nameContext) Date_formula_literal() IDate_formula_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_literalContext)
}

func (s *Alias_nameContext) Date_formula_n_literal_name() IDate_formula_n_literal_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_n_literal_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_n_literal_nameContext)
}

func (s *Alias_nameContext) Function_name() IFunction_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *Alias_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Alias_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Alias_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterAlias_name(s)
	}
}

func (s *Alias_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitAlias_name(s)
	}
}

func (s *Alias_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitAlias_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Alias_name() (localctx IAlias_nameContext) {
	localctx = NewAlias_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SOQLParserRULE_alias_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(202)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(197)
			p.Match(SOQLParserID)
		}

	case SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(198)
			p.Keywords_alias_allowed()
		}

	case SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(199)
			p.Date_formula_literal()
		}

	case SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(200)
			p.Date_formula_n_literal_name()
		}

	case SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR, SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM, SOQLParserDISTANCE, SOQLParserGEOLOCATION, SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(201)
			p.Function_name()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Alias_name() IAlias_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAlias_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAlias_nameContext)
}

func (s *AliasContext) AS() antlr.TerminalNode {
	return s.GetToken(SOQLParserAS, 0)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SOQLParserRULE_alias)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserAS {
		{
			p.SetState(204)
			p.Match(SOQLParserAS)
		}

	}
	{
		p.SetState(207)
		p.Alias_name()
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) Date_formula_literal() IDate_formula_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_literalContext)
}

func (s *LiteralContext) Date_formula_n_literal() IDate_formula_n_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_n_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_n_literalContext)
}

func (s *LiteralContext) Datetime_literal() IDatetime_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDatetime_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDatetime_literalContext)
}

func (s *LiteralContext) Date_literal() IDate_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_literalContext)
}

func (s *LiteralContext) Integer_literal() IInteger_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInteger_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInteger_literalContext)
}

func (s *LiteralContext) Real_literal() IReal_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReal_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReal_literalContext)
}

func (s *LiteralContext) String_literal() IString_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IString_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IString_literalContext)
}

func (s *LiteralContext) Boolean_literal() IBoolean_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoolean_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoolean_literalContext)
}

func (s *LiteralContext) Null_literal() INull_literalContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INull_literalContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INull_literalContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SOQLParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(218)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(209)
			p.Date_formula_literal()
		}

	case SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(210)
			p.Date_formula_n_literal()
		}

	case SOQLParserDATETIME:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(211)
			p.Datetime_literal()
		}

	case SOQLParserDATE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(212)
			p.Date_literal()
		}

	case SOQLParserUNSIGNED_INTEGER, SOQLParserSIGNED_INTEGER:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(213)
			p.Integer_literal()
		}

	case SOQLParserREAL_NUMBER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(214)
			p.Real_literal()
		}

	case SOQLParserSTRING_VALUE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(215)
			p.String_literal()
		}

	case SOQLParserFALSE, SOQLParserTRUE:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(216)
			p.Boolean_literal()
		}

	case SOQLParserNULL:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(217)
			p.Null_literal()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDate_formula_literalContext is an interface to support dynamic dispatch.
type IDate_formula_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_formula_literalContext differentiates from other interfaces.
	IsDate_formula_literalContext()
}

type Date_formula_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_formula_literalContext() *Date_formula_literalContext {
	var p = new(Date_formula_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_date_formula_literal
	return p
}

func (*Date_formula_literalContext) IsDate_formula_literalContext() {}

func NewDate_formula_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_formula_literalContext {
	var p = new(Date_formula_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_date_formula_literal

	return p
}

func (s *Date_formula_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_formula_literalContext) YESTERDAY() antlr.TerminalNode {
	return s.GetToken(SOQLParserYESTERDAY, 0)
}

func (s *Date_formula_literalContext) TODAY() antlr.TerminalNode {
	return s.GetToken(SOQLParserTODAY, 0)
}

func (s *Date_formula_literalContext) TOMORROW() antlr.TerminalNode {
	return s.GetToken(SOQLParserTOMORROW, 0)
}

func (s *Date_formula_literalContext) LAST_WEEK() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_WEEK, 0)
}

func (s *Date_formula_literalContext) THIS_WEEK() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_WEEK, 0)
}

func (s *Date_formula_literalContext) NEXT_WEEK() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_WEEK, 0)
}

func (s *Date_formula_literalContext) LAST_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_MONTH, 0)
}

func (s *Date_formula_literalContext) THIS_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_MONTH, 0)
}

func (s *Date_formula_literalContext) NEXT_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_MONTH, 0)
}

func (s *Date_formula_literalContext) LAST_90_DAYS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_90_DAYS, 0)
}

func (s *Date_formula_literalContext) NEXT_90_DAYS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_90_DAYS, 0)
}

func (s *Date_formula_literalContext) THIS_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_QUARTER, 0)
}

func (s *Date_formula_literalContext) LAST_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_QUARTER, 0)
}

func (s *Date_formula_literalContext) NEXT_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_QUARTER, 0)
}

func (s *Date_formula_literalContext) THIS_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_YEAR, 0)
}

func (s *Date_formula_literalContext) LAST_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_YEAR, 0)
}

func (s *Date_formula_literalContext) NEXT_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_YEAR, 0)
}

func (s *Date_formula_literalContext) THIS_FISCAL_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_FISCAL_QUARTER, 0)
}

func (s *Date_formula_literalContext) LAST_FISCAL_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_FISCAL_QUARTER, 0)
}

func (s *Date_formula_literalContext) NEXT_FISCAL_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_FISCAL_QUARTER, 0)
}

func (s *Date_formula_literalContext) THIS_FISCAL_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHIS_FISCAL_YEAR, 0)
}

func (s *Date_formula_literalContext) LAST_FISCAL_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_FISCAL_YEAR, 0)
}

func (s *Date_formula_literalContext) NEXT_FISCAL_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_FISCAL_YEAR, 0)
}

func (s *Date_formula_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_formula_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_formula_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterDate_formula_literal(s)
	}
}

func (s *Date_formula_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitDate_formula_literal(s)
	}
}

func (s *Date_formula_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitDate_formula_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Date_formula_literal() (localctx IDate_formula_literalContext) {
	localctx = NewDate_formula_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SOQLParserRULE_date_formula_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(220)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-77)&-(0x1f+1)) == 0 && ((1<<uint((_la-77)))&((1<<(SOQLParserYESTERDAY-77))|(1<<(SOQLParserTODAY-77))|(1<<(SOQLParserTOMORROW-77))|(1<<(SOQLParserLAST_WEEK-77))|(1<<(SOQLParserTHIS_WEEK-77))|(1<<(SOQLParserNEXT_WEEK-77))|(1<<(SOQLParserLAST_MONTH-77))|(1<<(SOQLParserTHIS_MONTH-77))|(1<<(SOQLParserNEXT_MONTH-77))|(1<<(SOQLParserLAST_90_DAYS-77))|(1<<(SOQLParserNEXT_90_DAYS-77))|(1<<(SOQLParserTHIS_QUARTER-77))|(1<<(SOQLParserLAST_QUARTER-77))|(1<<(SOQLParserNEXT_QUARTER-77))|(1<<(SOQLParserTHIS_YEAR-77))|(1<<(SOQLParserLAST_YEAR-77))|(1<<(SOQLParserNEXT_YEAR-77))|(1<<(SOQLParserTHIS_FISCAL_QUARTER-77))|(1<<(SOQLParserLAST_FISCAL_QUARTER-77))|(1<<(SOQLParserNEXT_FISCAL_QUARTER-77))|(1<<(SOQLParserTHIS_FISCAL_YEAR-77))|(1<<(SOQLParserLAST_FISCAL_YEAR-77))|(1<<(SOQLParserNEXT_FISCAL_YEAR-77)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDate_formula_n_literal_nameContext is an interface to support dynamic dispatch.
type IDate_formula_n_literal_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_formula_n_literal_nameContext differentiates from other interfaces.
	IsDate_formula_n_literal_nameContext()
}

type Date_formula_n_literal_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_formula_n_literal_nameContext() *Date_formula_n_literal_nameContext {
	var p = new(Date_formula_n_literal_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_date_formula_n_literal_name
	return p
}

func (*Date_formula_n_literal_nameContext) IsDate_formula_n_literal_nameContext() {}

func NewDate_formula_n_literal_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_formula_n_literal_nameContext {
	var p = new(Date_formula_n_literal_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_date_formula_n_literal_name

	return p
}

func (s *Date_formula_n_literal_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_formula_n_literal_nameContext) NEXT_N_DAYS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_DAYS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_DAYS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_DAYS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_DAYS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_DAYS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_WEEKS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_WEEKS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_WEEKS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_WEEKS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_WEEKS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_WEEKS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_MONTHS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_MONTHS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_MONTHS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_MONTHS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_MONTHS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_MONTHS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_QUARTERS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_QUARTERS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_QUARTERS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_QUARTERS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_QUARTERS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_QUARTERS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_YEARS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_YEARS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_YEARS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_YEARS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_YEARS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_YEARS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_FISCAL_QUARTERS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_FISCAL_QUARTERS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_FISCAL_QUARTERS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_FISCAL_QUARTERS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_FISCAL_QUARTERS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_FISCAL_QUARTERS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) NEXT_N_FISCAL_YEARS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNEXT_N_FISCAL_YEARS, 0)
}

func (s *Date_formula_n_literal_nameContext) LAST_N_FISCAL_YEARS() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST_N_FISCAL_YEARS, 0)
}

func (s *Date_formula_n_literal_nameContext) N_FISCAL_YEARS_AGO() antlr.TerminalNode {
	return s.GetToken(SOQLParserN_FISCAL_YEARS_AGO, 0)
}

func (s *Date_formula_n_literal_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_formula_n_literal_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_formula_n_literal_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterDate_formula_n_literal_name(s)
	}
}

func (s *Date_formula_n_literal_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitDate_formula_n_literal_name(s)
	}
}

func (s *Date_formula_n_literal_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitDate_formula_n_literal_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Date_formula_n_literal_name() (localctx IDate_formula_n_literal_nameContext) {
	localctx = NewDate_formula_n_literal_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SOQLParserRULE_date_formula_n_literal_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-100)&-(0x1f+1)) == 0 && ((1<<uint((_la-100)))&((1<<(SOQLParserNEXT_N_DAYS-100))|(1<<(SOQLParserLAST_N_DAYS-100))|(1<<(SOQLParserN_DAYS_AGO-100))|(1<<(SOQLParserNEXT_N_WEEKS-100))|(1<<(SOQLParserLAST_N_WEEKS-100))|(1<<(SOQLParserN_WEEKS_AGO-100))|(1<<(SOQLParserNEXT_N_MONTHS-100))|(1<<(SOQLParserLAST_N_MONTHS-100))|(1<<(SOQLParserN_MONTHS_AGO-100))|(1<<(SOQLParserNEXT_N_QUARTERS-100))|(1<<(SOQLParserLAST_N_QUARTERS-100))|(1<<(SOQLParserN_QUARTERS_AGO-100))|(1<<(SOQLParserNEXT_N_YEARS-100))|(1<<(SOQLParserLAST_N_YEARS-100))|(1<<(SOQLParserN_YEARS_AGO-100))|(1<<(SOQLParserNEXT_N_FISCAL_QUARTERS-100))|(1<<(SOQLParserLAST_N_FISCAL_QUARTERS-100))|(1<<(SOQLParserN_FISCAL_QUARTERS_AGO-100))|(1<<(SOQLParserNEXT_N_FISCAL_YEARS-100))|(1<<(SOQLParserLAST_N_FISCAL_YEARS-100))|(1<<(SOQLParserN_FISCAL_YEARS_AGO-100)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDate_formula_n_literalContext is an interface to support dynamic dispatch.
type IDate_formula_n_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_formula_n_literalContext differentiates from other interfaces.
	IsDate_formula_n_literalContext()
}

type Date_formula_n_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_formula_n_literalContext() *Date_formula_n_literalContext {
	var p = new(Date_formula_n_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_date_formula_n_literal
	return p
}

func (*Date_formula_n_literalContext) IsDate_formula_n_literalContext() {}

func NewDate_formula_n_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_formula_n_literalContext {
	var p = new(Date_formula_n_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_date_formula_n_literal

	return p
}

func (s *Date_formula_n_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_formula_n_literalContext) Date_formula_n_literal_name() IDate_formula_n_literal_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDate_formula_n_literal_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDate_formula_n_literal_nameContext)
}

func (s *Date_formula_n_literalContext) COLON() antlr.TerminalNode {
	return s.GetToken(SOQLParserCOLON, 0)
}

func (s *Date_formula_n_literalContext) UNSIGNED_INTEGER() antlr.TerminalNode {
	return s.GetToken(SOQLParserUNSIGNED_INTEGER, 0)
}

func (s *Date_formula_n_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_formula_n_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_formula_n_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterDate_formula_n_literal(s)
	}
}

func (s *Date_formula_n_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitDate_formula_n_literal(s)
	}
}

func (s *Date_formula_n_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitDate_formula_n_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Date_formula_n_literal() (localctx IDate_formula_n_literalContext) {
	localctx = NewDate_formula_n_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SOQLParserRULE_date_formula_n_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Date_formula_n_literal_name()
	}
	{
		p.SetState(225)
		p.Match(SOQLParserCOLON)
	}
	{
		p.SetState(226)
		p.Match(SOQLParserUNSIGNED_INTEGER)
	}

	return localctx
}

// IDatetime_literalContext is an interface to support dynamic dispatch.
type IDatetime_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatetime_literalContext differentiates from other interfaces.
	IsDatetime_literalContext()
}

type Datetime_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatetime_literalContext() *Datetime_literalContext {
	var p = new(Datetime_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_datetime_literal
	return p
}

func (*Datetime_literalContext) IsDatetime_literalContext() {}

func NewDatetime_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Datetime_literalContext {
	var p = new(Datetime_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_datetime_literal

	return p
}

func (s *Datetime_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Datetime_literalContext) DATETIME() antlr.TerminalNode {
	return s.GetToken(SOQLParserDATETIME, 0)
}

func (s *Datetime_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Datetime_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Datetime_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterDatetime_literal(s)
	}
}

func (s *Datetime_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitDatetime_literal(s)
	}
}

func (s *Datetime_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitDatetime_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Datetime_literal() (localctx IDatetime_literalContext) {
	localctx = NewDatetime_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SOQLParserRULE_datetime_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(SOQLParserDATETIME)
	}

	return localctx
}

// IDate_literalContext is an interface to support dynamic dispatch.
type IDate_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDate_literalContext differentiates from other interfaces.
	IsDate_literalContext()
}

type Date_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDate_literalContext() *Date_literalContext {
	var p = new(Date_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_date_literal
	return p
}

func (*Date_literalContext) IsDate_literalContext() {}

func NewDate_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Date_literalContext {
	var p = new(Date_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_date_literal

	return p
}

func (s *Date_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Date_literalContext) DATE() antlr.TerminalNode {
	return s.GetToken(SOQLParserDATE, 0)
}

func (s *Date_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Date_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Date_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterDate_literal(s)
	}
}

func (s *Date_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitDate_literal(s)
	}
}

func (s *Date_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitDate_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Date_literal() (localctx IDate_literalContext) {
	localctx = NewDate_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SOQLParserRULE_date_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.Match(SOQLParserDATE)
	}

	return localctx
}

// IInteger_literalContext is an interface to support dynamic dispatch.
type IInteger_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInteger_literalContext differentiates from other interfaces.
	IsInteger_literalContext()
}

type Integer_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInteger_literalContext() *Integer_literalContext {
	var p = new(Integer_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_integer_literal
	return p
}

func (*Integer_literalContext) IsInteger_literalContext() {}

func NewInteger_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Integer_literalContext {
	var p = new(Integer_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_integer_literal

	return p
}

func (s *Integer_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Integer_literalContext) SIGNED_INTEGER() antlr.TerminalNode {
	return s.GetToken(SOQLParserSIGNED_INTEGER, 0)
}

func (s *Integer_literalContext) UNSIGNED_INTEGER() antlr.TerminalNode {
	return s.GetToken(SOQLParserUNSIGNED_INTEGER, 0)
}

func (s *Integer_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Integer_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Integer_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterInteger_literal(s)
	}
}

func (s *Integer_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitInteger_literal(s)
	}
}

func (s *Integer_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitInteger_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Integer_literal() (localctx IInteger_literalContext) {
	localctx = NewInteger_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SOQLParserRULE_integer_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(232)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserUNSIGNED_INTEGER || _la == SOQLParserSIGNED_INTEGER) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IReal_literalContext is an interface to support dynamic dispatch.
type IReal_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsReal_literalContext differentiates from other interfaces.
	IsReal_literalContext()
}

type Real_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReal_literalContext() *Real_literalContext {
	var p = new(Real_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_real_literal
	return p
}

func (*Real_literalContext) IsReal_literalContext() {}

func NewReal_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Real_literalContext {
	var p = new(Real_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_real_literal

	return p
}

func (s *Real_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Real_literalContext) REAL_NUMBER() antlr.TerminalNode {
	return s.GetToken(SOQLParserREAL_NUMBER, 0)
}

func (s *Real_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Real_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Real_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterReal_literal(s)
	}
}

func (s *Real_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitReal_literal(s)
	}
}

func (s *Real_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitReal_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Real_literal() (localctx IReal_literalContext) {
	localctx = NewReal_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SOQLParserRULE_real_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(SOQLParserREAL_NUMBER)
	}

	return localctx
}

// IString_literalContext is an interface to support dynamic dispatch.
type IString_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsString_literalContext differentiates from other interfaces.
	IsString_literalContext()
}

type String_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_literalContext() *String_literalContext {
	var p = new(String_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_string_literal
	return p
}

func (*String_literalContext) IsString_literalContext() {}

func NewString_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_literalContext {
	var p = new(String_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_string_literal

	return p
}

func (s *String_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *String_literalContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(SOQLParserSTRING_VALUE, 0)
}

func (s *String_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterString_literal(s)
	}
}

func (s *String_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitString_literal(s)
	}
}

func (s *String_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitString_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) String_literal() (localctx IString_literalContext) {
	localctx = NewString_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SOQLParserRULE_string_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(SOQLParserSTRING_VALUE)
	}

	return localctx
}

// IBoolean_literalContext is an interface to support dynamic dispatch.
type IBoolean_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolean_literalContext differentiates from other interfaces.
	IsBoolean_literalContext()
}

type Boolean_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolean_literalContext() *Boolean_literalContext {
	var p = new(Boolean_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_boolean_literal
	return p
}

func (*Boolean_literalContext) IsBoolean_literalContext() {}

func NewBoolean_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Boolean_literalContext {
	var p = new(Boolean_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_boolean_literal

	return p
}

func (s *Boolean_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Boolean_literalContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SOQLParserTRUE, 0)
}

func (s *Boolean_literalContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SOQLParserFALSE, 0)
}

func (s *Boolean_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Boolean_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Boolean_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterBoolean_literal(s)
	}
}

func (s *Boolean_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitBoolean_literal(s)
	}
}

func (s *Boolean_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitBoolean_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Boolean_literal() (localctx IBoolean_literalContext) {
	localctx = NewBoolean_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SOQLParserRULE_boolean_literal)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserFALSE || _la == SOQLParserTRUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// INull_literalContext is an interface to support dynamic dispatch.
type INull_literalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNull_literalContext differentiates from other interfaces.
	IsNull_literalContext()
}

type Null_literalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNull_literalContext() *Null_literalContext {
	var p = new(Null_literalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_null_literal
	return p
}

func (*Null_literalContext) IsNull_literalContext() {}

func NewNull_literalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Null_literalContext {
	var p = new(Null_literalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_null_literal

	return p
}

func (s *Null_literalContext) GetParser() antlr.Parser { return s.parser }

func (s *Null_literalContext) NULL() antlr.TerminalNode {
	return s.GetToken(SOQLParserNULL, 0)
}

func (s *Null_literalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Null_literalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Null_literalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterNull_literal(s)
	}
}

func (s *Null_literalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitNull_literal(s)
	}
}

func (s *Null_literalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitNull_literal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Null_literal() (localctx INull_literalContext) {
	localctx = NewNull_literalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SOQLParserRULE_null_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(SOQLParserNULL)
	}

	return localctx
}

// IFunction_nameContext is an interface to support dynamic dispatch.
type IFunction_nameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_nameContext differentiates from other interfaces.
	IsFunction_nameContext()
}

type Function_nameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_nameContext() *Function_nameContext {
	var p = new(Function_nameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_name
	return p
}

func (*Function_nameContext) IsFunction_nameContext() {}

func NewFunction_nameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_nameContext {
	var p = new(Function_nameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_name

	return p
}

func (s *Function_nameContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_nameContext) Function_date() IFunction_dateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_dateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_dateContext)
}

func (s *Function_nameContext) Function_aggregate() IFunction_aggregateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_aggregateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_aggregateContext)
}

func (s *Function_nameContext) Function_location() IFunction_locationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_locationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_locationContext)
}

func (s *Function_nameContext) Function_other() IFunction_otherContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_otherContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_otherContext)
}

func (s *Function_nameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_nameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_nameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_name(s)
	}
}

func (s *Function_nameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_name(s)
	}
}

func (s *Function_nameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_name(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_name() (localctx IFunction_nameContext) {
	localctx = NewFunction_nameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SOQLParserRULE_function_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(246)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Function_date()
		}

	case SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(243)
			p.Function_aggregate()
		}

	case SOQLParserDISTANCE, SOQLParserGEOLOCATION:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(244)
			p.Function_location()
		}

	case SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(245)
			p.Function_other()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunction_dateContext is an interface to support dynamic dispatch.
type IFunction_dateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_dateContext differentiates from other interfaces.
	IsFunction_dateContext()
}

type Function_dateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_dateContext() *Function_dateContext {
	var p = new(Function_dateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_date
	return p
}

func (*Function_dateContext) IsFunction_dateContext() {}

func NewFunction_dateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_dateContext {
	var p = new(Function_dateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_date

	return p
}

func (s *Function_dateContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_dateContext) CALENDAR_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserCALENDAR_MONTH, 0)
}

func (s *Function_dateContext) CALENDAR_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserCALENDAR_QUARTER, 0)
}

func (s *Function_dateContext) CALENDAR_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserCALENDAR_YEAR, 0)
}

func (s *Function_dateContext) DAY_IN_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserDAY_IN_MONTH, 0)
}

func (s *Function_dateContext) DAY_IN_WEEK() antlr.TerminalNode {
	return s.GetToken(SOQLParserDAY_IN_WEEK, 0)
}

func (s *Function_dateContext) DAY_IN_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserDAY_IN_YEAR, 0)
}

func (s *Function_dateContext) DAY_ONLY() antlr.TerminalNode {
	return s.GetToken(SOQLParserDAY_ONLY, 0)
}

func (s *Function_dateContext) FISCAL_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserFISCAL_MONTH, 0)
}

func (s *Function_dateContext) FISCAL_QUARTER() antlr.TerminalNode {
	return s.GetToken(SOQLParserFISCAL_QUARTER, 0)
}

func (s *Function_dateContext) FISCAL_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserFISCAL_YEAR, 0)
}

func (s *Function_dateContext) HOUR_IN_DAY() antlr.TerminalNode {
	return s.GetToken(SOQLParserHOUR_IN_DAY, 0)
}

func (s *Function_dateContext) WEEK_IN_MONTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserWEEK_IN_MONTH, 0)
}

func (s *Function_dateContext) WEEK_IN_YEAR() antlr.TerminalNode {
	return s.GetToken(SOQLParserWEEK_IN_YEAR, 0)
}

func (s *Function_dateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_dateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_dateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_date(s)
	}
}

func (s *Function_dateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_date(s)
	}
}

func (s *Function_dateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_date(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_date() (localctx IFunction_dateContext) {
	localctx = NewFunction_dateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SOQLParserRULE_function_date)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-51)&-(0x1f+1)) == 0 && ((1<<uint((_la-51)))&((1<<(SOQLParserCALENDAR_MONTH-51))|(1<<(SOQLParserCALENDAR_QUARTER-51))|(1<<(SOQLParserCALENDAR_YEAR-51))|(1<<(SOQLParserDAY_IN_MONTH-51))|(1<<(SOQLParserDAY_IN_WEEK-51))|(1<<(SOQLParserDAY_IN_YEAR-51))|(1<<(SOQLParserDAY_ONLY-51))|(1<<(SOQLParserFISCAL_MONTH-51))|(1<<(SOQLParserFISCAL_QUARTER-51))|(1<<(SOQLParserFISCAL_YEAR-51))|(1<<(SOQLParserHOUR_IN_DAY-51))|(1<<(SOQLParserWEEK_IN_MONTH-51))|(1<<(SOQLParserWEEK_IN_YEAR-51)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_aggregateContext is an interface to support dynamic dispatch.
type IFunction_aggregateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_aggregateContext differentiates from other interfaces.
	IsFunction_aggregateContext()
}

type Function_aggregateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_aggregateContext() *Function_aggregateContext {
	var p = new(Function_aggregateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_aggregate
	return p
}

func (*Function_aggregateContext) IsFunction_aggregateContext() {}

func NewFunction_aggregateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_aggregateContext {
	var p = new(Function_aggregateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_aggregate

	return p
}

func (s *Function_aggregateContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_aggregateContext) AVG() antlr.TerminalNode {
	return s.GetToken(SOQLParserAVG, 0)
}

func (s *Function_aggregateContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SOQLParserCOUNT, 0)
}

func (s *Function_aggregateContext) COUNT_DISTINCT() antlr.TerminalNode {
	return s.GetToken(SOQLParserCOUNT_DISTINCT, 0)
}

func (s *Function_aggregateContext) MIN() antlr.TerminalNode {
	return s.GetToken(SOQLParserMIN, 0)
}

func (s *Function_aggregateContext) MAX() antlr.TerminalNode {
	return s.GetToken(SOQLParserMAX, 0)
}

func (s *Function_aggregateContext) SUM() antlr.TerminalNode {
	return s.GetToken(SOQLParserSUM, 0)
}

func (s *Function_aggregateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_aggregateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_aggregateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_aggregate(s)
	}
}

func (s *Function_aggregateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_aggregate(s)
	}
}

func (s *Function_aggregateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_aggregate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_aggregate() (localctx IFunction_aggregateContext) {
	localctx = NewFunction_aggregateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SOQLParserRULE_function_aggregate)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(SOQLParserAVG-64))|(1<<(SOQLParserCOUNT-64))|(1<<(SOQLParserCOUNT_DISTINCT-64))|(1<<(SOQLParserMIN-64))|(1<<(SOQLParserMAX-64))|(1<<(SOQLParserSUM-64)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_locationContext is an interface to support dynamic dispatch.
type IFunction_locationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_locationContext differentiates from other interfaces.
	IsFunction_locationContext()
}

type Function_locationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_locationContext() *Function_locationContext {
	var p = new(Function_locationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_location
	return p
}

func (*Function_locationContext) IsFunction_locationContext() {}

func NewFunction_locationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_locationContext {
	var p = new(Function_locationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_location

	return p
}

func (s *Function_locationContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_locationContext) DISTANCE() antlr.TerminalNode {
	return s.GetToken(SOQLParserDISTANCE, 0)
}

func (s *Function_locationContext) GEOLOCATION() antlr.TerminalNode {
	return s.GetToken(SOQLParserGEOLOCATION, 0)
}

func (s *Function_locationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_locationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_locationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_location(s)
	}
}

func (s *Function_locationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_location(s)
	}
}

func (s *Function_locationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_location(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_location() (localctx IFunction_locationContext) {
	localctx = NewFunction_locationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SOQLParserRULE_function_location)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserDISTANCE || _la == SOQLParserGEOLOCATION) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IFunction_otherContext is an interface to support dynamic dispatch.
type IFunction_otherContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_otherContext differentiates from other interfaces.
	IsFunction_otherContext()
}

type Function_otherContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_otherContext() *Function_otherContext {
	var p = new(Function_otherContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_other
	return p
}

func (*Function_otherContext) IsFunction_otherContext() {}

func NewFunction_otherContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_otherContext {
	var p = new(Function_otherContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_other

	return p
}

func (s *Function_otherContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_otherContext) FORMAT() antlr.TerminalNode {
	return s.GetToken(SOQLParserFORMAT, 0)
}

func (s *Function_otherContext) TOLABEL() antlr.TerminalNode {
	return s.GetToken(SOQLParserTOLABEL, 0)
}

func (s *Function_otherContext) CONVERT_TIME_ZONE() antlr.TerminalNode {
	return s.GetToken(SOQLParserCONVERT_TIME_ZONE, 0)
}

func (s *Function_otherContext) CONVERT_CURRENCY() antlr.TerminalNode {
	return s.GetToken(SOQLParserCONVERT_CURRENCY, 0)
}

func (s *Function_otherContext) GROUPING() antlr.TerminalNode {
	return s.GetToken(SOQLParserGROUPING, 0)
}

func (s *Function_otherContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_otherContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_otherContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_other(s)
	}
}

func (s *Function_otherContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_other(s)
	}
}

func (s *Function_otherContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_other(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_other() (localctx IFunction_otherContext) {
	localctx = NewFunction_otherContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SOQLParserRULE_function_other)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-72)&-(0x1f+1)) == 0 && ((1<<uint((_la-72)))&((1<<(SOQLParserFORMAT-72))|(1<<(SOQLParserTOLABEL-72))|(1<<(SOQLParserCONVERT_TIME_ZONE-72))|(1<<(SOQLParserCONVERT_CURRENCY-72))|(1<<(SOQLParserGROUPING-72)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISoql_queryContext is an interface to support dynamic dispatch.
type ISoql_queryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoql_queryContext differentiates from other interfaces.
	IsSoql_queryContext()
}

type Soql_queryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoql_queryContext() *Soql_queryContext {
	var p = new(Soql_queryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_soql_query
	return p
}

func (*Soql_queryContext) IsSoql_queryContext() {}

func NewSoql_queryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Soql_queryContext {
	var p = new(Soql_queryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_soql_query

	return p
}

func (s *Soql_queryContext) GetParser() antlr.Parser { return s.parser }

func (s *Soql_queryContext) Select_clause() ISelect_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelect_clauseContext)
}

func (s *Soql_queryContext) From_clause() IFrom_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_clauseContext)
}

func (s *Soql_queryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SOQLParserEOF, 0)
}

func (s *Soql_queryContext) Using_clause() IUsing_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsing_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUsing_clauseContext)
}

func (s *Soql_queryContext) Where_clause() IWhere_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *Soql_queryContext) With_clause() IWith_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWith_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWith_clauseContext)
}

func (s *Soql_queryContext) Groupby_clause() IGroupby_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupby_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupby_clauseContext)
}

func (s *Soql_queryContext) Orderby_clause() IOrderby_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderby_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderby_clauseContext)
}

func (s *Soql_queryContext) Limit_clause() ILimit_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimit_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimit_clauseContext)
}

func (s *Soql_queryContext) Offset_clause() IOffset_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffset_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOffset_clauseContext)
}

func (s *Soql_queryContext) For_clause() IFor_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_clauseContext)
}

func (s *Soql_queryContext) Update_clause() IUpdate_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_clauseContext)
}

func (s *Soql_queryContext) Having_clause() IHaving_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHaving_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHaving_clauseContext)
}

func (s *Soql_queryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Soql_queryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Soql_queryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSoql_query(s)
	}
}

func (s *Soql_queryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSoql_query(s)
	}
}

func (s *Soql_queryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSoql_query(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Soql_query() (localctx ISoql_queryContext) {
	localctx = NewSoql_queryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SOQLParserRULE_soql_query)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Select_clause()
	}
	{
		p.SetState(257)
		p.From_clause()
	}
	p.SetState(259)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserUSING {
		{
			p.SetState(258)
			p.Using_clause()
		}

	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserWHERE {
		{
			p.SetState(261)
			p.Where_clause()
		}

	}
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserWITH {
		{
			p.SetState(264)
			p.With_clause()
		}

	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserGROUP {
		{
			p.SetState(267)
			p.Groupby_clause()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SOQLParserHAVING {
			{
				p.SetState(268)
				p.Having_clause()
			}

		}

	}
	p.SetState(274)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserORDER {
		{
			p.SetState(273)
			p.Orderby_clause()
		}

	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserLIMIT {
		{
			p.SetState(276)
			p.Limit_clause()
		}

	}
	p.SetState(280)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserOFFSET {
		{
			p.SetState(279)
			p.Offset_clause()
		}

	}
	p.SetState(283)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserFOR {
		{
			p.SetState(282)
			p.For_clause()
		}

	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserUPDATE {
		{
			p.SetState(285)
			p.Update_clause()
		}

	}
	{
		p.SetState(288)
		p.Match(SOQLParserEOF)
	}

	return localctx
}

// ISelect_clauseContext is an interface to support dynamic dispatch.
type ISelect_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_clauseContext differentiates from other interfaces.
	IsSelect_clauseContext()
}

type Select_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_clauseContext() *Select_clauseContext {
	var p = new(Select_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_select_clause
	return p
}

func (*Select_clauseContext) IsSelect_clauseContext() {}

func NewSelect_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_clauseContext {
	var p = new(Select_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_select_clause

	return p
}

func (s *Select_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_clauseContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SOQLParserSELECT, 0)
}

func (s *Select_clauseContext) AllSelect_spec() []ISelect_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelect_specContext)(nil)).Elem())
	var tst = make([]ISelect_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelect_specContext)
		}
	}

	return tst
}

func (s *Select_clauseContext) Select_spec(i int) ISelect_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelect_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelect_specContext)
}

func (s *Select_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Select_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Select_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSelect_clause(s)
	}
}

func (s *Select_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSelect_clause(s)
	}
}

func (s *Select_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSelect_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Select_clause() (localctx ISelect_clauseContext) {
	localctx = NewSelect_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SOQLParserRULE_select_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(290)
		p.Match(SOQLParserSELECT)
	}
	{
		p.SetState(291)
		p.Select_spec()
	}
	p.SetState(296)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(292)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(293)
			p.Select_spec()
		}

		p.SetState(298)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFrom_clauseContext is an interface to support dynamic dispatch.
type IFrom_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFrom_clauseContext differentiates from other interfaces.
	IsFrom_clauseContext()
}

type From_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFrom_clauseContext() *From_clauseContext {
	var p = new(From_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_from_clause
	return p
}

func (*From_clauseContext) IsFrom_clauseContext() {}

func NewFrom_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *From_clauseContext {
	var p = new(From_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_from_clause

	return p
}

func (s *From_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *From_clauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(SOQLParserFROM, 0)
}

func (s *From_clauseContext) AllObject_spec() []IObject_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_specContext)(nil)).Elem())
	var tst = make([]IObject_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_specContext)
		}
	}

	return tst
}

func (s *From_clauseContext) Object_spec(i int) IObject_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObject_specContext)
}

func (s *From_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *From_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *From_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *From_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *From_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFrom_clause(s)
	}
}

func (s *From_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFrom_clause(s)
	}
}

func (s *From_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFrom_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) From_clause() (localctx IFrom_clauseContext) {
	localctx = NewFrom_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SOQLParserRULE_from_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(SOQLParserFROM)
	}
	{
		p.SetState(300)
		p.Object_spec()
	}
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(301)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(302)
			p.Object_spec()
		}

		p.SetState(307)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IUsing_clauseContext is an interface to support dynamic dispatch.
type IUsing_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUsing_clauseContext differentiates from other interfaces.
	IsUsing_clauseContext()
}

type Using_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUsing_clauseContext() *Using_clauseContext {
	var p = new(Using_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_using_clause
	return p
}

func (*Using_clauseContext) IsUsing_clauseContext() {}

func NewUsing_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Using_clauseContext {
	var p = new(Using_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_using_clause

	return p
}

func (s *Using_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Using_clauseContext) USING() antlr.TerminalNode {
	return s.GetToken(SOQLParserUSING, 0)
}

func (s *Using_clauseContext) SCOPE() antlr.TerminalNode {
	return s.GetToken(SOQLParserSCOPE, 0)
}

func (s *Using_clauseContext) Filter_scope_name() IFilter_scope_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFilter_scope_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFilter_scope_nameContext)
}

func (s *Using_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Using_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Using_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterUsing_clause(s)
	}
}

func (s *Using_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitUsing_clause(s)
	}
}

func (s *Using_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitUsing_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Using_clause() (localctx IUsing_clauseContext) {
	localctx = NewUsing_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SOQLParserRULE_using_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(SOQLParserUSING)
	}
	{
		p.SetState(309)
		p.Match(SOQLParserSCOPE)
	}
	{
		p.SetState(310)
		p.Filter_scope_name()
	}

	return localctx
}

// IWhere_clauseContext is an interface to support dynamic dispatch.
type IWhere_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhere_clauseContext differentiates from other interfaces.
	IsWhere_clauseContext()
}

type Where_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhere_clauseContext() *Where_clauseContext {
	var p = new(Where_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_where_clause
	return p
}

func (*Where_clauseContext) IsWhere_clauseContext() {}

func NewWhere_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Where_clauseContext {
	var p = new(Where_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_where_clause

	return p
}

func (s *Where_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Where_clauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SOQLParserWHERE, 0)
}

func (s *Where_clauseContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Where_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Where_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Where_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterWhere_clause(s)
	}
}

func (s *Where_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitWhere_clause(s)
	}
}

func (s *Where_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitWhere_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Where_clause() (localctx IWhere_clauseContext) {
	localctx = NewWhere_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SOQLParserRULE_where_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(SOQLParserWHERE)
	}
	{
		p.SetState(313)
		p.Condition()
	}

	return localctx
}

// IGroupby_clauseContext is an interface to support dynamic dispatch.
type IGroupby_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupby_clauseContext differentiates from other interfaces.
	IsGroupby_clauseContext()
}

type Groupby_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupby_clauseContext() *Groupby_clauseContext {
	var p = new(Groupby_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_groupby_clause
	return p
}

func (*Groupby_clauseContext) IsGroupby_clauseContext() {}

func NewGroupby_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Groupby_clauseContext {
	var p = new(Groupby_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_groupby_clause

	return p
}

func (s *Groupby_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Groupby_clauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SOQLParserGROUP, 0)
}

func (s *Groupby_clauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SOQLParserBY, 0)
}

func (s *Groupby_clauseContext) Group_by_rollup_clause() IGroup_by_rollup_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_rollup_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_rollup_clauseContext)
}

func (s *Groupby_clauseContext) Group_by_cube_clause() IGroup_by_cube_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_cube_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_cube_clauseContext)
}

func (s *Groupby_clauseContext) Group_by_plain_clause() IGroup_by_plain_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_plain_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_plain_clauseContext)
}

func (s *Groupby_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Groupby_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Groupby_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroupby_clause(s)
	}
}

func (s *Groupby_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroupby_clause(s)
	}
}

func (s *Groupby_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroupby_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Groupby_clause() (localctx IGroupby_clauseContext) {
	localctx = NewGroupby_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SOQLParserRULE_groupby_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(315)
		p.Match(SOQLParserGROUP)
	}
	{
		p.SetState(316)
		p.Match(SOQLParserBY)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserROLLUP:
		{
			p.SetState(317)
			p.Group_by_rollup_clause()
		}

	case SOQLParserCUBE:
		{
			p.SetState(318)
			p.Group_by_cube_clause()
		}

	case SOQLParserGROUP, SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN, SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR, SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM, SOQLParserDISTANCE, SOQLParserGEOLOCATION, SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING, SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR, SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO, SOQLParserID:
		{
			p.SetState(319)
			p.Group_by_plain_clause()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHaving_clauseContext is an interface to support dynamic dispatch.
type IHaving_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHaving_clauseContext differentiates from other interfaces.
	IsHaving_clauseContext()
}

type Having_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHaving_clauseContext() *Having_clauseContext {
	var p = new(Having_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_having_clause
	return p
}

func (*Having_clauseContext) IsHaving_clauseContext() {}

func NewHaving_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Having_clauseContext {
	var p = new(Having_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_having_clause

	return p
}

func (s *Having_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Having_clauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SOQLParserHAVING, 0)
}

func (s *Having_clauseContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *Having_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Having_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Having_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterHaving_clause(s)
	}
}

func (s *Having_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitHaving_clause(s)
	}
}

func (s *Having_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitHaving_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Having_clause() (localctx IHaving_clauseContext) {
	localctx = NewHaving_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SOQLParserRULE_having_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(SOQLParserHAVING)
	}
	{
		p.SetState(323)
		p.Condition()
	}

	return localctx
}

// IOrderby_clauseContext is an interface to support dynamic dispatch.
type IOrderby_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderby_clauseContext differentiates from other interfaces.
	IsOrderby_clauseContext()
}

type Orderby_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderby_clauseContext() *Orderby_clauseContext {
	var p = new(Orderby_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_orderby_clause
	return p
}

func (*Orderby_clauseContext) IsOrderby_clauseContext() {}

func NewOrderby_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Orderby_clauseContext {
	var p = new(Orderby_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_orderby_clause

	return p
}

func (s *Orderby_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Orderby_clauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SOQLParserORDER, 0)
}

func (s *Orderby_clauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SOQLParserBY, 0)
}

func (s *Orderby_clauseContext) Order_by_list() IOrder_by_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrder_by_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrder_by_listContext)
}

func (s *Orderby_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Orderby_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Orderby_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrderby_clause(s)
	}
}

func (s *Orderby_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrderby_clause(s)
	}
}

func (s *Orderby_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrderby_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Orderby_clause() (localctx IOrderby_clauseContext) {
	localctx = NewOrderby_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SOQLParserRULE_orderby_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Match(SOQLParserORDER)
	}
	{
		p.SetState(326)
		p.Match(SOQLParserBY)
	}
	{
		p.SetState(327)
		p.Order_by_list()
	}

	return localctx
}

// ILimit_clauseContext is an interface to support dynamic dispatch.
type ILimit_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimit_clauseContext differentiates from other interfaces.
	IsLimit_clauseContext()
}

type Limit_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimit_clauseContext() *Limit_clauseContext {
	var p = new(Limit_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_limit_clause
	return p
}

func (*Limit_clauseContext) IsLimit_clauseContext() {}

func NewLimit_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Limit_clauseContext {
	var p = new(Limit_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_limit_clause

	return p
}

func (s *Limit_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Limit_clauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SOQLParserLIMIT, 0)
}

func (s *Limit_clauseContext) UNSIGNED_INTEGER() antlr.TerminalNode {
	return s.GetToken(SOQLParserUNSIGNED_INTEGER, 0)
}

func (s *Limit_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Limit_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Limit_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterLimit_clause(s)
	}
}

func (s *Limit_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitLimit_clause(s)
	}
}

func (s *Limit_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitLimit_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Limit_clause() (localctx ILimit_clauseContext) {
	localctx = NewLimit_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SOQLParserRULE_limit_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(329)
		p.Match(SOQLParserLIMIT)
	}
	{
		p.SetState(330)
		p.Match(SOQLParserUNSIGNED_INTEGER)
	}

	return localctx
}

// IOffset_clauseContext is an interface to support dynamic dispatch.
type IOffset_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffset_clauseContext differentiates from other interfaces.
	IsOffset_clauseContext()
}

type Offset_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffset_clauseContext() *Offset_clauseContext {
	var p = new(Offset_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_offset_clause
	return p
}

func (*Offset_clauseContext) IsOffset_clauseContext() {}

func NewOffset_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Offset_clauseContext {
	var p = new(Offset_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_offset_clause

	return p
}

func (s *Offset_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Offset_clauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SOQLParserOFFSET, 0)
}

func (s *Offset_clauseContext) UNSIGNED_INTEGER() antlr.TerminalNode {
	return s.GetToken(SOQLParserUNSIGNED_INTEGER, 0)
}

func (s *Offset_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Offset_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Offset_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOffset_clause(s)
	}
}

func (s *Offset_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOffset_clause(s)
	}
}

func (s *Offset_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOffset_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Offset_clause() (localctx IOffset_clauseContext) {
	localctx = NewOffset_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SOQLParserRULE_offset_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(SOQLParserOFFSET)
	}
	{
		p.SetState(333)
		p.Match(SOQLParserUNSIGNED_INTEGER)
	}

	return localctx
}

// IFor_clauseContext is an interface to support dynamic dispatch.
type IFor_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_clauseContext differentiates from other interfaces.
	IsFor_clauseContext()
}

type For_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_clauseContext() *For_clauseContext {
	var p = new(For_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_for_clause
	return p
}

func (*For_clauseContext) IsFor_clauseContext() {}

func NewFor_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_clauseContext {
	var p = new(For_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_for_clause

	return p
}

func (s *For_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *For_clauseContext) FOR() antlr.TerminalNode {
	return s.GetToken(SOQLParserFOR, 0)
}

func (s *For_clauseContext) AllFor_value() []IFor_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFor_valueContext)(nil)).Elem())
	var tst = make([]IFor_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFor_valueContext)
		}
	}

	return tst
}

func (s *For_clauseContext) For_value(i int) IFor_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFor_valueContext)
}

func (s *For_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *For_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *For_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFor_clause(s)
	}
}

func (s *For_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFor_clause(s)
	}
}

func (s *For_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFor_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) For_clause() (localctx IFor_clauseContext) {
	localctx = NewFor_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SOQLParserRULE_for_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(335)
		p.Match(SOQLParserFOR)
	}
	{
		p.SetState(336)
		p.For_value()
	}
	p.SetState(339)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(337)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(338)
			p.For_value()
		}

	}
	p.SetState(343)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserCOMMA {
		{
			p.SetState(341)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(342)
			p.For_value()
		}

	}

	return localctx
}

// IUpdate_clauseContext is an interface to support dynamic dispatch.
type IUpdate_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_clauseContext differentiates from other interfaces.
	IsUpdate_clauseContext()
}

type Update_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_clauseContext() *Update_clauseContext {
	var p = new(Update_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_update_clause
	return p
}

func (*Update_clauseContext) IsUpdate_clauseContext() {}

func NewUpdate_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_clauseContext {
	var p = new(Update_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_update_clause

	return p
}

func (s *Update_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_clauseContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(SOQLParserUPDATE, 0)
}

func (s *Update_clauseContext) AllUpdate_value() []IUpdate_valueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUpdate_valueContext)(nil)).Elem())
	var tst = make([]IUpdate_valueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUpdate_valueContext)
		}
	}

	return tst
}

func (s *Update_clauseContext) Update_value(i int) IUpdate_valueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_valueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUpdate_valueContext)
}

func (s *Update_clauseContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, 0)
}

func (s *Update_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterUpdate_clause(s)
	}
}

func (s *Update_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitUpdate_clause(s)
	}
}

func (s *Update_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitUpdate_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Update_clause() (localctx IUpdate_clauseContext) {
	localctx = NewUpdate_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SOQLParserRULE_update_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(345)
		p.Match(SOQLParserUPDATE)
	}
	{
		p.SetState(346)
		p.Update_value()
	}
	p.SetState(349)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserCOMMA {
		{
			p.SetState(347)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(348)
			p.Update_value()
		}

	}

	return localctx
}

// ISoql_subqueryContext is an interface to support dynamic dispatch.
type ISoql_subqueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSoql_subqueryContext differentiates from other interfaces.
	IsSoql_subqueryContext()
}

type Soql_subqueryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySoql_subqueryContext() *Soql_subqueryContext {
	var p = new(Soql_subqueryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_soql_subquery
	return p
}

func (*Soql_subqueryContext) IsSoql_subqueryContext() {}

func NewSoql_subqueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Soql_subqueryContext {
	var p = new(Soql_subqueryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_soql_subquery

	return p
}

func (s *Soql_subqueryContext) GetParser() antlr.Parser { return s.parser }

func (s *Soql_subqueryContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Soql_subqueryContext) Subquery_select_clause() ISubquery_select_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubquery_select_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISubquery_select_clauseContext)
}

func (s *Soql_subqueryContext) From_clause() IFrom_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFrom_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFrom_clauseContext)
}

func (s *Soql_subqueryContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Soql_subqueryContext) Using_clause() IUsing_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUsing_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUsing_clauseContext)
}

func (s *Soql_subqueryContext) Where_clause() IWhere_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhere_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhere_clauseContext)
}

func (s *Soql_subqueryContext) With_clause() IWith_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWith_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWith_clauseContext)
}

func (s *Soql_subqueryContext) Orderby_clause() IOrderby_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderby_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderby_clauseContext)
}

func (s *Soql_subqueryContext) Limit_clause() ILimit_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimit_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimit_clauseContext)
}

func (s *Soql_subqueryContext) Offset_clause() IOffset_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOffset_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOffset_clauseContext)
}

func (s *Soql_subqueryContext) For_clause() IFor_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFor_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFor_clauseContext)
}

func (s *Soql_subqueryContext) Update_clause() IUpdate_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUpdate_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUpdate_clauseContext)
}

func (s *Soql_subqueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Soql_subqueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Soql_subqueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSoql_subquery(s)
	}
}

func (s *Soql_subqueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSoql_subquery(s)
	}
}

func (s *Soql_subqueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSoql_subquery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Soql_subquery() (localctx ISoql_subqueryContext) {
	localctx = NewSoql_subqueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SOQLParserRULE_soql_subquery)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(SOQLParserLPAREN)
	}
	{
		p.SetState(352)
		p.Subquery_select_clause()
	}
	{
		p.SetState(353)
		p.From_clause()
	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserUSING {
		{
			p.SetState(354)
			p.Using_clause()
		}

	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserWHERE {
		{
			p.SetState(357)
			p.Where_clause()
		}

	}
	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserWITH {
		{
			p.SetState(360)
			p.With_clause()
		}

	}
	p.SetState(364)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserORDER {
		{
			p.SetState(363)
			p.Orderby_clause()
		}

	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserLIMIT {
		{
			p.SetState(366)
			p.Limit_clause()
		}

	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserOFFSET {
		{
			p.SetState(369)
			p.Offset_clause()
		}

	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserFOR {
		{
			p.SetState(372)
			p.For_clause()
		}

	}
	p.SetState(376)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserUPDATE {
		{
			p.SetState(375)
			p.Update_clause()
		}

	}
	{
		p.SetState(378)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// ISubquery_select_clauseContext is an interface to support dynamic dispatch.
type ISubquery_select_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubquery_select_clauseContext differentiates from other interfaces.
	IsSubquery_select_clauseContext()
}

type Subquery_select_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubquery_select_clauseContext() *Subquery_select_clauseContext {
	var p = new(Subquery_select_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_subquery_select_clause
	return p
}

func (*Subquery_select_clauseContext) IsSubquery_select_clauseContext() {}

func NewSubquery_select_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subquery_select_clauseContext {
	var p = new(Subquery_select_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_subquery_select_clause

	return p
}

func (s *Subquery_select_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Subquery_select_clauseContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SOQLParserSELECT, 0)
}

func (s *Subquery_select_clauseContext) AllSubquery_select_spec() []ISubquery_select_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISubquery_select_specContext)(nil)).Elem())
	var tst = make([]ISubquery_select_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISubquery_select_specContext)
		}
	}

	return tst
}

func (s *Subquery_select_clauseContext) Subquery_select_spec(i int) ISubquery_select_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISubquery_select_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISubquery_select_specContext)
}

func (s *Subquery_select_clauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Subquery_select_clauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Subquery_select_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subquery_select_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subquery_select_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSubquery_select_clause(s)
	}
}

func (s *Subquery_select_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSubquery_select_clause(s)
	}
}

func (s *Subquery_select_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSubquery_select_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Subquery_select_clause() (localctx ISubquery_select_clauseContext) {
	localctx = NewSubquery_select_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SOQLParserRULE_subquery_select_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(380)
		p.Match(SOQLParserSELECT)
	}
	{
		p.SetState(381)
		p.Subquery_select_spec()
	}
	p.SetState(386)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(382)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(383)
			p.Subquery_select_spec()
		}

		p.SetState(388)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelect_specContext is an interface to support dynamic dispatch.
type ISelect_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelect_specContext differentiates from other interfaces.
	IsSelect_specContext()
}

type Select_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelect_specContext() *Select_specContext {
	var p = new(Select_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_select_spec
	return p
}

func (*Select_specContext) IsSelect_specContext() {}

func NewSelect_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Select_specContext {
	var p = new(Select_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_select_spec

	return p
}

func (s *Select_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Select_specContext) Field_spec() IField_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specContext)
}

func (s *Select_specContext) Function_call_spec() IFunction_call_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_call_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_call_specContext)
}

func (s *Select_specContext) Soql_subquery() ISoql_subqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoql_subqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoql_subqueryContext)
}

func (s *Select_specContext) Typeof_spec() ITypeof_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeof_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeof_specContext)
}

func (s *Select_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Select_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Select_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSelect_spec(s)
	}
}

func (s *Select_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSelect_spec(s)
	}
}

func (s *Select_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSelect_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Select_spec() (localctx ISelect_specContext) {
	localctx = NewSelect_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SOQLParserRULE_select_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)
			p.Field_spec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)
			p.Function_call_spec()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(391)
			p.Soql_subquery()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(392)
			p.Typeof_spec()
		}

	}

	return localctx
}

// ISubquery_select_specContext is an interface to support dynamic dispatch.
type ISubquery_select_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSubquery_select_specContext differentiates from other interfaces.
	IsSubquery_select_specContext()
}

type Subquery_select_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySubquery_select_specContext() *Subquery_select_specContext {
	var p = new(Subquery_select_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_subquery_select_spec
	return p
}

func (*Subquery_select_specContext) IsSubquery_select_specContext() {}

func NewSubquery_select_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Subquery_select_specContext {
	var p = new(Subquery_select_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_subquery_select_spec

	return p
}

func (s *Subquery_select_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Subquery_select_specContext) Field_spec() IField_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_specContext)
}

func (s *Subquery_select_specContext) Function_call_spec() IFunction_call_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_call_specContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_call_specContext)
}

func (s *Subquery_select_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subquery_select_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Subquery_select_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSubquery_select_spec(s)
	}
}

func (s *Subquery_select_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSubquery_select_spec(s)
	}
}

func (s *Subquery_select_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSubquery_select_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Subquery_select_spec() (localctx ISubquery_select_specContext) {
	localctx = NewSubquery_select_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SOQLParserRULE_subquery_select_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(397)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(395)
			p.Field_spec()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Function_call_spec()
		}

	}

	return localctx
}

// IField_specContext is an interface to support dynamic dispatch.
type IField_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_specContext differentiates from other interfaces.
	IsField_specContext()
}

type Field_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_specContext() *Field_specContext {
	var p = new(Field_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_field_spec
	return p
}

func (*Field_specContext) IsField_specContext() {}

func NewField_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_specContext {
	var p = new(Field_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_field_spec

	return p
}

func (s *Field_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_specContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Field_specContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Field_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterField_spec(s)
	}
}

func (s *Field_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitField_spec(s)
	}
}

func (s *Field_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitField_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Field_spec() (localctx IField_specContext) {
	localctx = NewField_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SOQLParserRULE_field_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(399)
		p.Field()
	}
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserAS || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SOQLParserABOVE-34))|(1<<(SOQLParserABOVE_OR_BELOW-34))|(1<<(SOQLParserAT-34))|(1<<(SOQLParserBELOW-34))|(1<<(SOQLParserCATEGORY-34))|(1<<(SOQLParserDATA-34))|(1<<(SOQLParserEND-34))|(1<<(SOQLParserOFFSET-34))|(1<<(SOQLParserORDER-34))|(1<<(SOQLParserREFERENCE-34))|(1<<(SOQLParserSCOPE-34))|(1<<(SOQLParserTRACKING-34))|(1<<(SOQLParserTHEN-34))|(1<<(SOQLParserTYPEOF-34))|(1<<(SOQLParserVIEW-34))|(1<<(SOQLParserVIEWSTAT-34))|(1<<(SOQLParserWHEN-34))|(1<<(SOQLParserCALENDAR_MONTH-34))|(1<<(SOQLParserCALENDAR_QUARTER-34))|(1<<(SOQLParserCALENDAR_YEAR-34))|(1<<(SOQLParserDAY_IN_MONTH-34))|(1<<(SOQLParserDAY_IN_WEEK-34))|(1<<(SOQLParserDAY_IN_YEAR-34))|(1<<(SOQLParserDAY_ONLY-34))|(1<<(SOQLParserFISCAL_MONTH-34))|(1<<(SOQLParserFISCAL_QUARTER-34))|(1<<(SOQLParserFISCAL_YEAR-34))|(1<<(SOQLParserHOUR_IN_DAY-34))|(1<<(SOQLParserWEEK_IN_MONTH-34))|(1<<(SOQLParserWEEK_IN_YEAR-34))|(1<<(SOQLParserAVG-34))|(1<<(SOQLParserCOUNT-34)))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(SOQLParserCOUNT_DISTINCT-66))|(1<<(SOQLParserMIN-66))|(1<<(SOQLParserMAX-66))|(1<<(SOQLParserSUM-66))|(1<<(SOQLParserDISTANCE-66))|(1<<(SOQLParserGEOLOCATION-66))|(1<<(SOQLParserFORMAT-66))|(1<<(SOQLParserTOLABEL-66))|(1<<(SOQLParserCONVERT_TIME_ZONE-66))|(1<<(SOQLParserCONVERT_CURRENCY-66))|(1<<(SOQLParserGROUPING-66))|(1<<(SOQLParserYESTERDAY-66))|(1<<(SOQLParserTODAY-66))|(1<<(SOQLParserTOMORROW-66))|(1<<(SOQLParserLAST_WEEK-66))|(1<<(SOQLParserTHIS_WEEK-66))|(1<<(SOQLParserNEXT_WEEK-66))|(1<<(SOQLParserLAST_MONTH-66))|(1<<(SOQLParserTHIS_MONTH-66))|(1<<(SOQLParserNEXT_MONTH-66))|(1<<(SOQLParserLAST_90_DAYS-66))|(1<<(SOQLParserNEXT_90_DAYS-66))|(1<<(SOQLParserTHIS_QUARTER-66))|(1<<(SOQLParserLAST_QUARTER-66))|(1<<(SOQLParserNEXT_QUARTER-66))|(1<<(SOQLParserTHIS_YEAR-66))|(1<<(SOQLParserLAST_YEAR-66))|(1<<(SOQLParserNEXT_YEAR-66))|(1<<(SOQLParserTHIS_FISCAL_QUARTER-66))|(1<<(SOQLParserLAST_FISCAL_QUARTER-66))|(1<<(SOQLParserNEXT_FISCAL_QUARTER-66))|(1<<(SOQLParserTHIS_FISCAL_YEAR-66)))) != 0) || (((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(SOQLParserLAST_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_N_DAYS-98))|(1<<(SOQLParserLAST_N_DAYS-98))|(1<<(SOQLParserN_DAYS_AGO-98))|(1<<(SOQLParserNEXT_N_WEEKS-98))|(1<<(SOQLParserLAST_N_WEEKS-98))|(1<<(SOQLParserN_WEEKS_AGO-98))|(1<<(SOQLParserNEXT_N_MONTHS-98))|(1<<(SOQLParserLAST_N_MONTHS-98))|(1<<(SOQLParserN_MONTHS_AGO-98))|(1<<(SOQLParserNEXT_N_QUARTERS-98))|(1<<(SOQLParserLAST_N_QUARTERS-98))|(1<<(SOQLParserN_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_YEARS-98))|(1<<(SOQLParserLAST_N_YEARS-98))|(1<<(SOQLParserN_YEARS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserLAST_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserN_FISCAL_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_YEARS-98))|(1<<(SOQLParserLAST_N_FISCAL_YEARS-98))|(1<<(SOQLParserN_FISCAL_YEARS_AGO-98)))) != 0) || _la == SOQLParserID {
		{
			p.SetState(400)
			p.Alias()
		}

	}

	return localctx
}

// IFunction_call_specContext is an interface to support dynamic dispatch.
type IFunction_call_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_call_specContext differentiates from other interfaces.
	IsFunction_call_specContext()
}

type Function_call_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_call_specContext() *Function_call_specContext {
	var p = new(Function_call_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_call_spec
	return p
}

func (*Function_call_specContext) IsFunction_call_specContext() {}

func NewFunction_call_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_call_specContext {
	var p = new(Function_call_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_call_spec

	return p
}

func (s *Function_call_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_call_specContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Function_call_specContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Function_call_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_call_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_call_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_call_spec(s)
	}
}

func (s *Function_call_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_call_spec(s)
	}
}

func (s *Function_call_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_call_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_call_spec() (localctx IFunction_call_specContext) {
	localctx = NewFunction_call_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SOQLParserRULE_function_call_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Function_call()
	}
	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserAS || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SOQLParserABOVE-34))|(1<<(SOQLParserABOVE_OR_BELOW-34))|(1<<(SOQLParserAT-34))|(1<<(SOQLParserBELOW-34))|(1<<(SOQLParserCATEGORY-34))|(1<<(SOQLParserDATA-34))|(1<<(SOQLParserEND-34))|(1<<(SOQLParserOFFSET-34))|(1<<(SOQLParserORDER-34))|(1<<(SOQLParserREFERENCE-34))|(1<<(SOQLParserSCOPE-34))|(1<<(SOQLParserTRACKING-34))|(1<<(SOQLParserTHEN-34))|(1<<(SOQLParserTYPEOF-34))|(1<<(SOQLParserVIEW-34))|(1<<(SOQLParserVIEWSTAT-34))|(1<<(SOQLParserWHEN-34))|(1<<(SOQLParserCALENDAR_MONTH-34))|(1<<(SOQLParserCALENDAR_QUARTER-34))|(1<<(SOQLParserCALENDAR_YEAR-34))|(1<<(SOQLParserDAY_IN_MONTH-34))|(1<<(SOQLParserDAY_IN_WEEK-34))|(1<<(SOQLParserDAY_IN_YEAR-34))|(1<<(SOQLParserDAY_ONLY-34))|(1<<(SOQLParserFISCAL_MONTH-34))|(1<<(SOQLParserFISCAL_QUARTER-34))|(1<<(SOQLParserFISCAL_YEAR-34))|(1<<(SOQLParserHOUR_IN_DAY-34))|(1<<(SOQLParserWEEK_IN_MONTH-34))|(1<<(SOQLParserWEEK_IN_YEAR-34))|(1<<(SOQLParserAVG-34))|(1<<(SOQLParserCOUNT-34)))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(SOQLParserCOUNT_DISTINCT-66))|(1<<(SOQLParserMIN-66))|(1<<(SOQLParserMAX-66))|(1<<(SOQLParserSUM-66))|(1<<(SOQLParserDISTANCE-66))|(1<<(SOQLParserGEOLOCATION-66))|(1<<(SOQLParserFORMAT-66))|(1<<(SOQLParserTOLABEL-66))|(1<<(SOQLParserCONVERT_TIME_ZONE-66))|(1<<(SOQLParserCONVERT_CURRENCY-66))|(1<<(SOQLParserGROUPING-66))|(1<<(SOQLParserYESTERDAY-66))|(1<<(SOQLParserTODAY-66))|(1<<(SOQLParserTOMORROW-66))|(1<<(SOQLParserLAST_WEEK-66))|(1<<(SOQLParserTHIS_WEEK-66))|(1<<(SOQLParserNEXT_WEEK-66))|(1<<(SOQLParserLAST_MONTH-66))|(1<<(SOQLParserTHIS_MONTH-66))|(1<<(SOQLParserNEXT_MONTH-66))|(1<<(SOQLParserLAST_90_DAYS-66))|(1<<(SOQLParserNEXT_90_DAYS-66))|(1<<(SOQLParserTHIS_QUARTER-66))|(1<<(SOQLParserLAST_QUARTER-66))|(1<<(SOQLParserNEXT_QUARTER-66))|(1<<(SOQLParserTHIS_YEAR-66))|(1<<(SOQLParserLAST_YEAR-66))|(1<<(SOQLParserNEXT_YEAR-66))|(1<<(SOQLParserTHIS_FISCAL_QUARTER-66))|(1<<(SOQLParserLAST_FISCAL_QUARTER-66))|(1<<(SOQLParserNEXT_FISCAL_QUARTER-66))|(1<<(SOQLParserTHIS_FISCAL_YEAR-66)))) != 0) || (((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(SOQLParserLAST_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_N_DAYS-98))|(1<<(SOQLParserLAST_N_DAYS-98))|(1<<(SOQLParserN_DAYS_AGO-98))|(1<<(SOQLParserNEXT_N_WEEKS-98))|(1<<(SOQLParserLAST_N_WEEKS-98))|(1<<(SOQLParserN_WEEKS_AGO-98))|(1<<(SOQLParserNEXT_N_MONTHS-98))|(1<<(SOQLParserLAST_N_MONTHS-98))|(1<<(SOQLParserN_MONTHS_AGO-98))|(1<<(SOQLParserNEXT_N_QUARTERS-98))|(1<<(SOQLParserLAST_N_QUARTERS-98))|(1<<(SOQLParserN_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_YEARS-98))|(1<<(SOQLParserLAST_N_YEARS-98))|(1<<(SOQLParserN_YEARS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserLAST_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserN_FISCAL_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_YEARS-98))|(1<<(SOQLParserLAST_N_FISCAL_YEARS-98))|(1<<(SOQLParserN_FISCAL_YEARS_AGO-98)))) != 0) || _la == SOQLParserID {
		{
			p.SetState(404)
			p.Alias()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Field_name() IField_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_nameContext)
}

func (s *FieldContext) Object_prefix() IObject_prefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_prefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_prefixContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SOQLParserRULE_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(408)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(407)
			p.Object_prefix()
		}

	}
	{
		p.SetState(410)
		p.Field_name()
	}

	return localctx
}

// IFunction_callContext is an interface to support dynamic dispatch.
type IFunction_callContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_callContext differentiates from other interfaces.
	IsFunction_callContext()
}

type Function_callContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_callContext() *Function_callContext {
	var p = new(Function_callContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_call
	return p
}

func (*Function_callContext) IsFunction_callContext() {}

func NewFunction_callContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_callContext {
	var p = new(Function_callContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_call

	return p
}

func (s *Function_callContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_callContext) Function_name() IFunction_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_nameContext)
}

func (s *Function_callContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Function_callContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Function_callContext) Function_parameter_list() IFunction_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_parameter_listContext)
}

func (s *Function_callContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_callContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_callContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_call(s)
	}
}

func (s *Function_callContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_call(s)
	}
}

func (s *Function_callContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_call(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_call() (localctx IFunction_callContext) {
	localctx = NewFunction_callContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SOQLParserRULE_function_call)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(412)
		p.Function_name()
	}
	{
		p.SetState(413)
		p.Match(SOQLParserLPAREN)
	}
	p.SetState(415)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SOQLParserSTRING_VALUE)|(1<<SOQLParserFALSE)|(1<<SOQLParserGROUP)|(1<<SOQLParserNULL)|(1<<SOQLParserTRUE))) != 0) || (((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SOQLParserABOVE-34))|(1<<(SOQLParserABOVE_OR_BELOW-34))|(1<<(SOQLParserAT-34))|(1<<(SOQLParserBELOW-34))|(1<<(SOQLParserCATEGORY-34))|(1<<(SOQLParserDATA-34))|(1<<(SOQLParserEND-34))|(1<<(SOQLParserOFFSET-34))|(1<<(SOQLParserORDER-34))|(1<<(SOQLParserREFERENCE-34))|(1<<(SOQLParserSCOPE-34))|(1<<(SOQLParserTRACKING-34))|(1<<(SOQLParserTHEN-34))|(1<<(SOQLParserTYPEOF-34))|(1<<(SOQLParserVIEW-34))|(1<<(SOQLParserVIEWSTAT-34))|(1<<(SOQLParserWHEN-34))|(1<<(SOQLParserCALENDAR_MONTH-34))|(1<<(SOQLParserCALENDAR_QUARTER-34))|(1<<(SOQLParserCALENDAR_YEAR-34))|(1<<(SOQLParserDAY_IN_MONTH-34))|(1<<(SOQLParserDAY_IN_WEEK-34))|(1<<(SOQLParserDAY_IN_YEAR-34))|(1<<(SOQLParserDAY_ONLY-34))|(1<<(SOQLParserFISCAL_MONTH-34))|(1<<(SOQLParserFISCAL_QUARTER-34))|(1<<(SOQLParserFISCAL_YEAR-34))|(1<<(SOQLParserHOUR_IN_DAY-34))|(1<<(SOQLParserWEEK_IN_MONTH-34))|(1<<(SOQLParserWEEK_IN_YEAR-34))|(1<<(SOQLParserAVG-34))|(1<<(SOQLParserCOUNT-34)))) != 0) || (((_la-66)&-(0x1f+1)) == 0 && ((1<<uint((_la-66)))&((1<<(SOQLParserCOUNT_DISTINCT-66))|(1<<(SOQLParserMIN-66))|(1<<(SOQLParserMAX-66))|(1<<(SOQLParserSUM-66))|(1<<(SOQLParserDISTANCE-66))|(1<<(SOQLParserGEOLOCATION-66))|(1<<(SOQLParserFORMAT-66))|(1<<(SOQLParserTOLABEL-66))|(1<<(SOQLParserCONVERT_TIME_ZONE-66))|(1<<(SOQLParserCONVERT_CURRENCY-66))|(1<<(SOQLParserGROUPING-66))|(1<<(SOQLParserYESTERDAY-66))|(1<<(SOQLParserTODAY-66))|(1<<(SOQLParserTOMORROW-66))|(1<<(SOQLParserLAST_WEEK-66))|(1<<(SOQLParserTHIS_WEEK-66))|(1<<(SOQLParserNEXT_WEEK-66))|(1<<(SOQLParserLAST_MONTH-66))|(1<<(SOQLParserTHIS_MONTH-66))|(1<<(SOQLParserNEXT_MONTH-66))|(1<<(SOQLParserLAST_90_DAYS-66))|(1<<(SOQLParserNEXT_90_DAYS-66))|(1<<(SOQLParserTHIS_QUARTER-66))|(1<<(SOQLParserLAST_QUARTER-66))|(1<<(SOQLParserNEXT_QUARTER-66))|(1<<(SOQLParserTHIS_YEAR-66))|(1<<(SOQLParserLAST_YEAR-66))|(1<<(SOQLParserNEXT_YEAR-66))|(1<<(SOQLParserTHIS_FISCAL_QUARTER-66))|(1<<(SOQLParserLAST_FISCAL_QUARTER-66))|(1<<(SOQLParserNEXT_FISCAL_QUARTER-66))|(1<<(SOQLParserTHIS_FISCAL_YEAR-66)))) != 0) || (((_la-98)&-(0x1f+1)) == 0 && ((1<<uint((_la-98)))&((1<<(SOQLParserLAST_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_FISCAL_YEAR-98))|(1<<(SOQLParserNEXT_N_DAYS-98))|(1<<(SOQLParserLAST_N_DAYS-98))|(1<<(SOQLParserN_DAYS_AGO-98))|(1<<(SOQLParserNEXT_N_WEEKS-98))|(1<<(SOQLParserLAST_N_WEEKS-98))|(1<<(SOQLParserN_WEEKS_AGO-98))|(1<<(SOQLParserNEXT_N_MONTHS-98))|(1<<(SOQLParserLAST_N_MONTHS-98))|(1<<(SOQLParserN_MONTHS_AGO-98))|(1<<(SOQLParserNEXT_N_QUARTERS-98))|(1<<(SOQLParserLAST_N_QUARTERS-98))|(1<<(SOQLParserN_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_YEARS-98))|(1<<(SOQLParserLAST_N_YEARS-98))|(1<<(SOQLParserN_YEARS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserLAST_N_FISCAL_QUARTERS-98))|(1<<(SOQLParserN_FISCAL_QUARTERS_AGO-98))|(1<<(SOQLParserNEXT_N_FISCAL_YEARS-98))|(1<<(SOQLParserLAST_N_FISCAL_YEARS-98))|(1<<(SOQLParserN_FISCAL_YEARS_AGO-98)))) != 0) || (((_la-136)&-(0x1f+1)) == 0 && ((1<<uint((_la-136)))&((1<<(SOQLParserID-136))|(1<<(SOQLParserDATE-136))|(1<<(SOQLParserDATETIME-136))|(1<<(SOQLParserUNSIGNED_INTEGER-136))|(1<<(SOQLParserREAL_NUMBER-136))|(1<<(SOQLParserSIGNED_INTEGER-136)))) != 0) {
		{
			p.SetState(414)
			p.Function_parameter_list()
		}

	}
	{
		p.SetState(417)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// IFunction_parameter_listContext is an interface to support dynamic dispatch.
type IFunction_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_parameter_listContext differentiates from other interfaces.
	IsFunction_parameter_listContext()
}

type Function_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_parameter_listContext() *Function_parameter_listContext {
	var p = new(Function_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_parameter_list
	return p
}

func (*Function_parameter_listContext) IsFunction_parameter_listContext() {}

func NewFunction_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_parameter_listContext {
	var p = new(Function_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_parameter_list

	return p
}

func (s *Function_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_parameter_listContext) AllFunction_parameter() []IFunction_parameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunction_parameterContext)(nil)).Elem())
	var tst = make([]IFunction_parameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunction_parameterContext)
		}
	}

	return tst
}

func (s *Function_parameter_listContext) Function_parameter(i int) IFunction_parameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_parameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunction_parameterContext)
}

func (s *Function_parameter_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Function_parameter_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Function_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_parameter_list(s)
	}
}

func (s *Function_parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_parameter_list(s)
	}
}

func (s *Function_parameter_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_parameter_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_parameter_list() (localctx IFunction_parameter_listContext) {
	localctx = NewFunction_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SOQLParserRULE_function_parameter_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Function_parameter()
	}
	p.SetState(424)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(420)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(421)
			p.Function_parameter()
		}

		p.SetState(426)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IFunction_parameterContext is an interface to support dynamic dispatch.
type IFunction_parameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunction_parameterContext differentiates from other interfaces.
	IsFunction_parameterContext()
}

type Function_parameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunction_parameterContext() *Function_parameterContext {
	var p = new(Function_parameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_function_parameter
	return p
}

func (*Function_parameterContext) IsFunction_parameterContext() {}

func NewFunction_parameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Function_parameterContext {
	var p = new(Function_parameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_function_parameter

	return p
}

func (s *Function_parameterContext) GetParser() antlr.Parser { return s.parser }

func (s *Function_parameterContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Function_parameterContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Function_parameterContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Function_parameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Function_parameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Function_parameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFunction_parameter(s)
	}
}

func (s *Function_parameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFunction_parameter(s)
	}
}

func (s *Function_parameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFunction_parameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Function_parameter() (localctx IFunction_parameterContext) {
	localctx = NewFunction_parameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SOQLParserRULE_function_parameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(427)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Literal()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(429)
			p.Function_call()
		}

	}

	return localctx
}

// ITypeof_specContext is an interface to support dynamic dispatch.
type ITypeof_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeof_specContext differentiates from other interfaces.
	IsTypeof_specContext()
}

type Typeof_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeof_specContext() *Typeof_specContext {
	var p = new(Typeof_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_typeof_spec
	return p
}

func (*Typeof_specContext) IsTypeof_specContext() {}

func NewTypeof_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typeof_specContext {
	var p = new(Typeof_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_typeof_spec

	return p
}

func (s *Typeof_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Typeof_specContext) TYPEOF() antlr.TerminalNode {
	return s.GetToken(SOQLParserTYPEOF, 0)
}

func (s *Typeof_specContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Typeof_specContext) Typeof_when_then_clause_list() ITypeof_when_then_clause_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeof_when_then_clause_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeof_when_then_clause_listContext)
}

func (s *Typeof_specContext) END() antlr.TerminalNode {
	return s.GetToken(SOQLParserEND, 0)
}

func (s *Typeof_specContext) Typeof_else_clause() ITypeof_else_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeof_else_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeof_else_clauseContext)
}

func (s *Typeof_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typeof_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typeof_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterTypeof_spec(s)
	}
}

func (s *Typeof_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitTypeof_spec(s)
	}
}

func (s *Typeof_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitTypeof_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Typeof_spec() (localctx ITypeof_specContext) {
	localctx = NewTypeof_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SOQLParserRULE_typeof_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(432)
		p.Match(SOQLParserTYPEOF)
	}
	{
		p.SetState(433)
		p.Field()
	}
	{
		p.SetState(434)
		p.Typeof_when_then_clause_list()
	}
	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserELSE {
		{
			p.SetState(435)
			p.Typeof_else_clause()
		}

	}
	{
		p.SetState(438)
		p.Match(SOQLParserEND)
	}

	return localctx
}

// ITypeof_when_then_clause_listContext is an interface to support dynamic dispatch.
type ITypeof_when_then_clause_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeof_when_then_clause_listContext differentiates from other interfaces.
	IsTypeof_when_then_clause_listContext()
}

type Typeof_when_then_clause_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeof_when_then_clause_listContext() *Typeof_when_then_clause_listContext {
	var p = new(Typeof_when_then_clause_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_typeof_when_then_clause_list
	return p
}

func (*Typeof_when_then_clause_listContext) IsTypeof_when_then_clause_listContext() {}

func NewTypeof_when_then_clause_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typeof_when_then_clause_listContext {
	var p = new(Typeof_when_then_clause_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_typeof_when_then_clause_list

	return p
}

func (s *Typeof_when_then_clause_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Typeof_when_then_clause_listContext) AllTypeof_when_then_clause() []ITypeof_when_then_clauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeof_when_then_clauseContext)(nil)).Elem())
	var tst = make([]ITypeof_when_then_clauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeof_when_then_clauseContext)
		}
	}

	return tst
}

func (s *Typeof_when_then_clause_listContext) Typeof_when_then_clause(i int) ITypeof_when_then_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeof_when_then_clauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeof_when_then_clauseContext)
}

func (s *Typeof_when_then_clause_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typeof_when_then_clause_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typeof_when_then_clause_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterTypeof_when_then_clause_list(s)
	}
}

func (s *Typeof_when_then_clause_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitTypeof_when_then_clause_list(s)
	}
}

func (s *Typeof_when_then_clause_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitTypeof_when_then_clause_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Typeof_when_then_clause_list() (localctx ITypeof_when_then_clause_listContext) {
	localctx = NewTypeof_when_then_clause_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SOQLParserRULE_typeof_when_then_clause_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == SOQLParserWHEN {
		{
			p.SetState(440)
			p.Typeof_when_then_clause()
		}

		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITypeof_when_then_clauseContext is an interface to support dynamic dispatch.
type ITypeof_when_then_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeof_when_then_clauseContext differentiates from other interfaces.
	IsTypeof_when_then_clauseContext()
}

type Typeof_when_then_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeof_when_then_clauseContext() *Typeof_when_then_clauseContext {
	var p = new(Typeof_when_then_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_typeof_when_then_clause
	return p
}

func (*Typeof_when_then_clauseContext) IsTypeof_when_then_clauseContext() {}

func NewTypeof_when_then_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typeof_when_then_clauseContext {
	var p = new(Typeof_when_then_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_typeof_when_then_clause

	return p
}

func (s *Typeof_when_then_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Typeof_when_then_clauseContext) WHEN() antlr.TerminalNode {
	return s.GetToken(SOQLParserWHEN, 0)
}

func (s *Typeof_when_then_clauseContext) Object_name() IObject_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_nameContext)
}

func (s *Typeof_when_then_clauseContext) Typeof_then_clause() ITypeof_then_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeof_then_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeof_then_clauseContext)
}

func (s *Typeof_when_then_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typeof_when_then_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typeof_when_then_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterTypeof_when_then_clause(s)
	}
}

func (s *Typeof_when_then_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitTypeof_when_then_clause(s)
	}
}

func (s *Typeof_when_then_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitTypeof_when_then_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Typeof_when_then_clause() (localctx ITypeof_when_then_clauseContext) {
	localctx = NewTypeof_when_then_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SOQLParserRULE_typeof_when_then_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Match(SOQLParserWHEN)
	}
	{
		p.SetState(446)
		p.Object_name()
	}
	{
		p.SetState(447)
		p.Typeof_then_clause()
	}

	return localctx
}

// ITypeof_then_clauseContext is an interface to support dynamic dispatch.
type ITypeof_then_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeof_then_clauseContext differentiates from other interfaces.
	IsTypeof_then_clauseContext()
}

type Typeof_then_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeof_then_clauseContext() *Typeof_then_clauseContext {
	var p = new(Typeof_then_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_typeof_then_clause
	return p
}

func (*Typeof_then_clauseContext) IsTypeof_then_clauseContext() {}

func NewTypeof_then_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typeof_then_clauseContext {
	var p = new(Typeof_then_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_typeof_then_clause

	return p
}

func (s *Typeof_then_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Typeof_then_clauseContext) THEN() antlr.TerminalNode {
	return s.GetToken(SOQLParserTHEN, 0)
}

func (s *Typeof_then_clauseContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Typeof_then_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typeof_then_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typeof_then_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterTypeof_then_clause(s)
	}
}

func (s *Typeof_then_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitTypeof_then_clause(s)
	}
}

func (s *Typeof_then_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitTypeof_then_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Typeof_then_clause() (localctx ITypeof_then_clauseContext) {
	localctx = NewTypeof_then_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SOQLParserRULE_typeof_then_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Match(SOQLParserTHEN)
	}
	{
		p.SetState(450)
		p.Field_list()
	}

	return localctx
}

// ITypeof_else_clauseContext is an interface to support dynamic dispatch.
type ITypeof_else_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeof_else_clauseContext differentiates from other interfaces.
	IsTypeof_else_clauseContext()
}

type Typeof_else_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeof_else_clauseContext() *Typeof_else_clauseContext {
	var p = new(Typeof_else_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_typeof_else_clause
	return p
}

func (*Typeof_else_clauseContext) IsTypeof_else_clauseContext() {}

func NewTypeof_else_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Typeof_else_clauseContext {
	var p = new(Typeof_else_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_typeof_else_clause

	return p
}

func (s *Typeof_else_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Typeof_else_clauseContext) ELSE() antlr.TerminalNode {
	return s.GetToken(SOQLParserELSE, 0)
}

func (s *Typeof_else_clauseContext) Field_list() IField_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_listContext)
}

func (s *Typeof_else_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Typeof_else_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Typeof_else_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterTypeof_else_clause(s)
	}
}

func (s *Typeof_else_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitTypeof_else_clause(s)
	}
}

func (s *Typeof_else_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitTypeof_else_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Typeof_else_clause() (localctx ITypeof_else_clauseContext) {
	localctx = NewTypeof_else_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SOQLParserRULE_typeof_else_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(SOQLParserELSE)
	}
	{
		p.SetState(453)
		p.Field_list()
	}

	return localctx
}

// IField_listContext is an interface to support dynamic dispatch.
type IField_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_listContext differentiates from other interfaces.
	IsField_listContext()
}

type Field_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_listContext() *Field_listContext {
	var p = new(Field_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_field_list
	return p
}

func (*Field_listContext) IsField_listContext() {}

func NewField_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_listContext {
	var p = new(Field_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_field_list

	return p
}

func (s *Field_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_listContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *Field_listContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Field_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Field_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Field_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterField_list(s)
	}
}

func (s *Field_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitField_list(s)
	}
}

func (s *Field_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitField_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Field_list() (localctx IField_listContext) {
	localctx = NewField_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SOQLParserRULE_field_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Field()
	}
	p.SetState(460)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(456)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(457)
			p.Field()
		}

		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IObject_specContext is an interface to support dynamic dispatch.
type IObject_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_specContext differentiates from other interfaces.
	IsObject_specContext()
}

type Object_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_specContext() *Object_specContext {
	var p = new(Object_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_object_spec
	return p
}

func (*Object_specContext) IsObject_specContext() {}

func NewObject_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_specContext {
	var p = new(Object_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_object_spec

	return p
}

func (s *Object_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_specContext) Object_name() IObject_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_nameContext)
}

func (s *Object_specContext) Object_prefix() IObject_prefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_prefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObject_prefixContext)
}

func (s *Object_specContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *Object_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterObject_spec(s)
	}
}

func (s *Object_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitObject_spec(s)
	}
}

func (s *Object_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitObject_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Object_spec() (localctx IObject_specContext) {
	localctx = NewObject_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SOQLParserRULE_object_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(464)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(463)
			p.Object_prefix()
		}

	}
	{
		p.SetState(466)
		p.Object_name()
	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(467)
			p.Alias()
		}

	}

	return localctx
}

// IObject_prefixContext is an interface to support dynamic dispatch.
type IObject_prefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObject_prefixContext differentiates from other interfaces.
	IsObject_prefixContext()
}

type Object_prefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_prefixContext() *Object_prefixContext {
	var p = new(Object_prefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_object_prefix
	return p
}

func (*Object_prefixContext) IsObject_prefixContext() {}

func NewObject_prefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_prefixContext {
	var p = new(Object_prefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_object_prefix

	return p
}

func (s *Object_prefixContext) GetParser() antlr.Parser { return s.parser }

func (s *Object_prefixContext) AllObject_name() []IObject_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObject_nameContext)(nil)).Elem())
	var tst = make([]IObject_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObject_nameContext)
		}
	}

	return tst
}

func (s *Object_prefixContext) Object_name(i int) IObject_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObject_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObject_nameContext)
}

func (s *Object_prefixContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserDOT)
}

func (s *Object_prefixContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserDOT, i)
}

func (s *Object_prefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_prefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_prefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterObject_prefix(s)
	}
}

func (s *Object_prefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitObject_prefix(s)
	}
}

func (s *Object_prefixContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitObject_prefix(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Object_prefix() (localctx IObject_prefixContext) {
	localctx = NewObject_prefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SOQLParserRULE_object_prefix)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(473)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(470)
				p.Object_name()
			}
			{
				p.SetState(471)
				p.Match(SOQLParserDOT)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(475)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext())
	}

	return localctx
}

// IComparison_operatorContext is an interface to support dynamic dispatch.
type IComparison_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparison_operatorContext differentiates from other interfaces.
	IsComparison_operatorContext()
}

type Comparison_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparison_operatorContext() *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_comparison_operator
	return p
}

func (*Comparison_operatorContext) IsComparison_operatorContext() {}

func NewComparison_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comparison_operatorContext {
	var p = new(Comparison_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_comparison_operator

	return p
}

func (s *Comparison_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Comparison_operatorContext) EQ() antlr.TerminalNode {
	return s.GetToken(SOQLParserEQ, 0)
}

func (s *Comparison_operatorContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(SOQLParserNOT_EQ, 0)
}

func (s *Comparison_operatorContext) LET() antlr.TerminalNode {
	return s.GetToken(SOQLParserLET, 0)
}

func (s *Comparison_operatorContext) GET() antlr.TerminalNode {
	return s.GetToken(SOQLParserGET, 0)
}

func (s *Comparison_operatorContext) GTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserGTH, 0)
}

func (s *Comparison_operatorContext) LTH() antlr.TerminalNode {
	return s.GetToken(SOQLParserLTH, 0)
}

func (s *Comparison_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comparison_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Comparison_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitComparison_operator(s)
	}
}

func (s *Comparison_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitComparison_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Comparison_operator() (localctx IComparison_operatorContext) {
	localctx = NewComparison_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SOQLParserRULE_comparison_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(487)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(477)
			p.Match(SOQLParserEQ)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(478)
			p.Match(SOQLParserNOT_EQ)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(479)
			p.Match(SOQLParserLET)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(480)
			p.Match(SOQLParserGET)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(481)
			p.Match(SOQLParserGTH)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(482)
			p.Match(SOQLParserLTH)
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(483)
			p.Match(SOQLParserGTH)
		}
		{
			p.SetState(484)
			p.Match(SOQLParserEQ)
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(485)
			p.Match(SOQLParserLTH)
		}
		{
			p.SetState(486)
			p.Match(SOQLParserEQ)
		}

	}

	return localctx
}

// ISet_operatorContext is an interface to support dynamic dispatch.
type ISet_operatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_operatorContext differentiates from other interfaces.
	IsSet_operatorContext()
}

type Set_operatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_operatorContext() *Set_operatorContext {
	var p = new(Set_operatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_set_operator
	return p
}

func (*Set_operatorContext) IsSet_operatorContext() {}

func NewSet_operatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_operatorContext {
	var p = new(Set_operatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_set_operator

	return p
}

func (s *Set_operatorContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_operatorContext) IN() antlr.TerminalNode {
	return s.GetToken(SOQLParserIN, 0)
}

func (s *Set_operatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(SOQLParserNOT, 0)
}

func (s *Set_operatorContext) INCLUDES() antlr.TerminalNode {
	return s.GetToken(SOQLParserINCLUDES, 0)
}

func (s *Set_operatorContext) EXCLUDES() antlr.TerminalNode {
	return s.GetToken(SOQLParserEXCLUDES, 0)
}

func (s *Set_operatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_operatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_operatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSet_operator(s)
	}
}

func (s *Set_operatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSet_operator(s)
	}
}

func (s *Set_operatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSet_operator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Set_operator() (localctx ISet_operatorContext) {
	localctx = NewSet_operatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SOQLParserRULE_set_operator)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(494)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserIN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(489)
			p.Match(SOQLParserIN)
		}

	case SOQLParserNOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(490)
			p.Match(SOQLParserNOT)
		}
		{
			p.SetState(491)
			p.Match(SOQLParserIN)
		}

	case SOQLParserINCLUDES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(492)
			p.Match(SOQLParserINCLUDES)
		}

	case SOQLParserEXCLUDES:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(493)
			p.Match(SOQLParserEXCLUDES)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConditionContext is an interface to support dynamic dispatch.
type IConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConditionContext differentiates from other interfaces.
	IsConditionContext()
}

type ConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionContext() *ConditionContext {
	var p = new(ConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_condition
	return p
}

func (*ConditionContext) IsConditionContext() {}

func NewConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionContext {
	var p = new(ConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_condition

	return p
}

func (s *ConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionContext) AllCondition1() []ICondition1Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICondition1Context)(nil)).Elem())
	var tst = make([]ICondition1Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICondition1Context)
		}
	}

	return tst
}

func (s *ConditionContext) Condition1(i int) ICondition1Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition1Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICondition1Context)
}

func (s *ConditionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserOR)
}

func (s *ConditionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserOR, i)
}

func (s *ConditionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserAND)
}

func (s *ConditionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserAND, i)
}

func (s *ConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterCondition(s)
	}
}

func (s *ConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitCondition(s)
	}
}

func (s *ConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Condition() (localctx IConditionContext) {
	localctx = NewConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SOQLParserRULE_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Condition1()
	}
	p.SetState(501)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserAND || _la == SOQLParserOR {
		{
			p.SetState(497)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SOQLParserAND || _la == SOQLParserOR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(498)
			p.Condition1()
		}

		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ICondition1Context is an interface to support dynamic dispatch.
type ICondition1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondition1Context differentiates from other interfaces.
	IsCondition1Context()
}

type Condition1Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition1Context() *Condition1Context {
	var p = new(Condition1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_condition1
	return p
}

func (*Condition1Context) IsCondition1Context() {}

func NewCondition1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition1Context {
	var p = new(Condition1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_condition1

	return p
}

func (s *Condition1Context) GetParser() antlr.Parser { return s.parser }

func (s *Condition1Context) Simple_condition() ISimple_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimple_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimple_conditionContext)
}

func (s *Condition1Context) Parenthesis() IParenthesisContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParenthesisContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParenthesisContext)
}

func (s *Condition1Context) NOT() antlr.TerminalNode {
	return s.GetToken(SOQLParserNOT, 0)
}

func (s *Condition1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterCondition1(s)
	}
}

func (s *Condition1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitCondition1(s)
	}
}

func (s *Condition1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitCondition1(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Condition1() (localctx ICondition1Context) {
	localctx = NewCondition1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SOQLParserRULE_condition1)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(505)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserNOT {
		{
			p.SetState(504)
			p.Match(SOQLParserNOT)
		}

	}
	p.SetState(509)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserGROUP, SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN, SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR, SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM, SOQLParserDISTANCE, SOQLParserGEOLOCATION, SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING, SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR, SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO, SOQLParserID:
		{
			p.SetState(507)
			p.Simple_condition()
		}

	case SOQLParserLPAREN:
		{
			p.SetState(508)
			p.Parenthesis()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IParenthesisContext is an interface to support dynamic dispatch.
type IParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParenthesisContext differentiates from other interfaces.
	IsParenthesisContext()
}

type ParenthesisContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParenthesisContext() *ParenthesisContext {
	var p = new(ParenthesisContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_parenthesis
	return p
}

func (*ParenthesisContext) IsParenthesisContext() {}

func NewParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_parenthesis

	return p
}

func (s *ParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *ParenthesisContext) Condition() IConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionContext)
}

func (s *ParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

func (s *ParenthesisContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitParenthesis(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Parenthesis() (localctx IParenthesisContext) {
	localctx = NewParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SOQLParserRULE_parenthesis)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)
		p.Match(SOQLParserLPAREN)
	}
	{
		p.SetState(512)
		p.Condition()
	}
	{
		p.SetState(513)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// ISimple_conditionContext is an interface to support dynamic dispatch.
type ISimple_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimple_conditionContext differentiates from other interfaces.
	IsSimple_conditionContext()
}

type Simple_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimple_conditionContext() *Simple_conditionContext {
	var p = new(Simple_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_simple_condition
	return p
}

func (*Simple_conditionContext) IsSimple_conditionContext() {}

func NewSimple_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Simple_conditionContext {
	var p = new(Simple_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_simple_condition

	return p
}

func (s *Simple_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Simple_conditionContext) Field_based_condition() IField_based_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_based_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_based_conditionContext)
}

func (s *Simple_conditionContext) Set_based_condition() ISet_based_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_based_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_based_conditionContext)
}

func (s *Simple_conditionContext) Like_based_condition() ILike_based_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILike_based_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILike_based_conditionContext)
}

func (s *Simple_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Simple_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Simple_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSimple_condition(s)
	}
}

func (s *Simple_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSimple_condition(s)
	}
}

func (s *Simple_conditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSimple_condition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Simple_condition() (localctx ISimple_conditionContext) {
	localctx = NewSimple_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SOQLParserRULE_simple_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(515)
			p.Field_based_condition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(516)
			p.Set_based_condition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(517)
			p.Like_based_condition()
		}

	}

	return localctx
}

// IField_based_conditionContext is an interface to support dynamic dispatch.
type IField_based_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsField_based_conditionContext differentiates from other interfaces.
	IsField_based_conditionContext()
}

type Field_based_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyField_based_conditionContext() *Field_based_conditionContext {
	var p = new(Field_based_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_field_based_condition
	return p
}

func (*Field_based_conditionContext) IsField_based_conditionContext() {}

func NewField_based_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Field_based_conditionContext {
	var p = new(Field_based_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_field_based_condition

	return p
}

func (s *Field_based_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Field_based_conditionContext) Condition_field() ICondition_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_fieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondition_fieldContext)
}

func (s *Field_based_conditionContext) Comparison_operator() IComparison_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparison_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparison_operatorContext)
}

func (s *Field_based_conditionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Field_based_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Field_based_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Field_based_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterField_based_condition(s)
	}
}

func (s *Field_based_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitField_based_condition(s)
	}
}

func (s *Field_based_conditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitField_based_condition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Field_based_condition() (localctx IField_based_conditionContext) {
	localctx = NewField_based_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SOQLParserRULE_field_based_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(520)
		p.Condition_field()
	}
	{
		p.SetState(521)
		p.Comparison_operator()
	}
	{
		p.SetState(522)
		p.Literal()
	}

	return localctx
}

// ISet_based_conditionContext is an interface to support dynamic dispatch.
type ISet_based_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_based_conditionContext differentiates from other interfaces.
	IsSet_based_conditionContext()
}

type Set_based_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_based_conditionContext() *Set_based_conditionContext {
	var p = new(Set_based_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_set_based_condition
	return p
}

func (*Set_based_conditionContext) IsSet_based_conditionContext() {}

func NewSet_based_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_based_conditionContext {
	var p = new(Set_based_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_set_based_condition

	return p
}

func (s *Set_based_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_based_conditionContext) Condition_field() ICondition_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_fieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondition_fieldContext)
}

func (s *Set_based_conditionContext) Set_operator() ISet_operatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_operatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_operatorContext)
}

func (s *Set_based_conditionContext) Soql_subquery() ISoql_subqueryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISoql_subqueryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISoql_subqueryContext)
}

func (s *Set_based_conditionContext) Set_values() ISet_valuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_valuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_valuesContext)
}

func (s *Set_based_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_based_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_based_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSet_based_condition(s)
	}
}

func (s *Set_based_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSet_based_condition(s)
	}
}

func (s *Set_based_conditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSet_based_condition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Set_based_condition() (localctx ISet_based_conditionContext) {
	localctx = NewSet_based_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SOQLParserRULE_set_based_condition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Condition_field()
	}
	{
		p.SetState(525)
		p.Set_operator()
	}
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(526)
			p.Soql_subquery()
		}

	case 2:
		{
			p.SetState(527)
			p.Set_values()
		}

	}

	return localctx
}

// ILike_based_conditionContext is an interface to support dynamic dispatch.
type ILike_based_conditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLike_based_conditionContext differentiates from other interfaces.
	IsLike_based_conditionContext()
}

type Like_based_conditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLike_based_conditionContext() *Like_based_conditionContext {
	var p = new(Like_based_conditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_like_based_condition
	return p
}

func (*Like_based_conditionContext) IsLike_based_conditionContext() {}

func NewLike_based_conditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Like_based_conditionContext {
	var p = new(Like_based_conditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_like_based_condition

	return p
}

func (s *Like_based_conditionContext) GetParser() antlr.Parser { return s.parser }

func (s *Like_based_conditionContext) Condition_field() ICondition_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICondition_fieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICondition_fieldContext)
}

func (s *Like_based_conditionContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SOQLParserLIKE, 0)
}

func (s *Like_based_conditionContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(SOQLParserSTRING_VALUE, 0)
}

func (s *Like_based_conditionContext) LIKE_STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(SOQLParserLIKE_STRING_VALUE, 0)
}

func (s *Like_based_conditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Like_based_conditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Like_based_conditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterLike_based_condition(s)
	}
}

func (s *Like_based_conditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitLike_based_condition(s)
	}
}

func (s *Like_based_conditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitLike_based_condition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Like_based_condition() (localctx ILike_based_conditionContext) {
	localctx = NewLike_based_conditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, SOQLParserRULE_like_based_condition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(530)
		p.Condition_field()
	}
	{
		p.SetState(531)
		p.Match(SOQLParserLIKE)
	}
	{
		p.SetState(532)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserSTRING_VALUE || _la == SOQLParserLIKE_STRING_VALUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ICondition_fieldContext is an interface to support dynamic dispatch.
type ICondition_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCondition_fieldContext differentiates from other interfaces.
	IsCondition_fieldContext()
}

type Condition_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCondition_fieldContext() *Condition_fieldContext {
	var p = new(Condition_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_condition_field
	return p
}

func (*Condition_fieldContext) IsCondition_fieldContext() {}

func NewCondition_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Condition_fieldContext {
	var p = new(Condition_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_condition_field

	return p
}

func (s *Condition_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Condition_fieldContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Condition_fieldContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Condition_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Condition_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Condition_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterCondition_field(s)
	}
}

func (s *Condition_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitCondition_field(s)
	}
}

func (s *Condition_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitCondition_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Condition_field() (localctx ICondition_fieldContext) {
	localctx = NewCondition_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, SOQLParserRULE_condition_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(536)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(534)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(535)
			p.Function_call()
		}

	}

	return localctx
}

// ISet_valuesContext is an interface to support dynamic dispatch.
type ISet_valuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_valuesContext differentiates from other interfaces.
	IsSet_valuesContext()
}

type Set_valuesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_valuesContext() *Set_valuesContext {
	var p = new(Set_valuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_set_values
	return p
}

func (*Set_valuesContext) IsSet_valuesContext() {}

func NewSet_valuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_valuesContext {
	var p = new(Set_valuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_set_values

	return p
}

func (s *Set_valuesContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_valuesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Set_valuesContext) Set_value_list() ISet_value_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_value_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_value_listContext)
}

func (s *Set_valuesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Set_valuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_valuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_valuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSet_values(s)
	}
}

func (s *Set_valuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSet_values(s)
	}
}

func (s *Set_valuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSet_values(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Set_values() (localctx ISet_valuesContext) {
	localctx = NewSet_valuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, SOQLParserRULE_set_values)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(538)
		p.Match(SOQLParserLPAREN)
	}
	{
		p.SetState(539)
		p.Set_value_list()
	}
	{
		p.SetState(540)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// ISet_value_listContext is an interface to support dynamic dispatch.
type ISet_value_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSet_value_listContext differentiates from other interfaces.
	IsSet_value_listContext()
}

type Set_value_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySet_value_listContext() *Set_value_listContext {
	var p = new(Set_value_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_set_value_list
	return p
}

func (*Set_value_listContext) IsSet_value_listContext() {}

func NewSet_value_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_value_listContext {
	var p = new(Set_value_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_set_value_list

	return p
}

func (s *Set_value_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_value_listContext) AllLiteral() []ILiteralContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILiteralContext)(nil)).Elem())
	var tst = make([]ILiteralContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILiteralContext)
		}
	}

	return tst
}

func (s *Set_value_listContext) Literal(i int) ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *Set_value_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Set_value_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Set_value_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_value_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_value_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterSet_value_list(s)
	}
}

func (s *Set_value_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitSet_value_list(s)
	}
}

func (s *Set_value_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitSet_value_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Set_value_list() (localctx ISet_value_listContext) {
	localctx = NewSet_value_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, SOQLParserRULE_set_value_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		p.Literal()
	}
	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(543)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(544)
			p.Literal()
		}

		p.SetState(549)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IWith_clauseContext is an interface to support dynamic dispatch.
type IWith_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWith_clauseContext differentiates from other interfaces.
	IsWith_clauseContext()
}

type With_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWith_clauseContext() *With_clauseContext {
	var p = new(With_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_with_clause
	return p
}

func (*With_clauseContext) IsWith_clauseContext() {}

func NewWith_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *With_clauseContext {
	var p = new(With_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_with_clause

	return p
}

func (s *With_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *With_clauseContext) WITH() antlr.TerminalNode {
	return s.GetToken(SOQLParserWITH, 0)
}

func (s *With_clauseContext) With_plain_clause() IWith_plain_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWith_plain_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWith_plain_clauseContext)
}

func (s *With_clauseContext) With_data_category_clause() IWith_data_category_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWith_data_category_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWith_data_category_clauseContext)
}

func (s *With_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *With_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *With_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterWith_clause(s)
	}
}

func (s *With_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitWith_clause(s)
	}
}

func (s *With_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitWith_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) With_clause() (localctx IWith_clauseContext) {
	localctx = NewWith_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, SOQLParserRULE_with_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(550)
		p.Match(SOQLParserWITH)
	}
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(551)
			p.With_plain_clause()
		}

	case 2:
		{
			p.SetState(552)
			p.With_data_category_clause()
		}

	}

	return localctx
}

// IWith_plain_clauseContext is an interface to support dynamic dispatch.
type IWith_plain_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWith_plain_clauseContext differentiates from other interfaces.
	IsWith_plain_clauseContext()
}

type With_plain_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWith_plain_clauseContext() *With_plain_clauseContext {
	var p = new(With_plain_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_with_plain_clause
	return p
}

func (*With_plain_clauseContext) IsWith_plain_clauseContext() {}

func NewWith_plain_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *With_plain_clauseContext {
	var p = new(With_plain_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_with_plain_clause

	return p
}

func (s *With_plain_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *With_plain_clauseContext) Field_based_condition() IField_based_conditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IField_based_conditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IField_based_conditionContext)
}

func (s *With_plain_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *With_plain_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *With_plain_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterWith_plain_clause(s)
	}
}

func (s *With_plain_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitWith_plain_clause(s)
	}
}

func (s *With_plain_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitWith_plain_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) With_plain_clause() (localctx IWith_plain_clauseContext) {
	localctx = NewWith_plain_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, SOQLParserRULE_with_plain_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(555)
		p.Field_based_condition()
	}

	return localctx
}

// IWith_data_category_clauseContext is an interface to support dynamic dispatch.
type IWith_data_category_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWith_data_category_clauseContext differentiates from other interfaces.
	IsWith_data_category_clauseContext()
}

type With_data_category_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWith_data_category_clauseContext() *With_data_category_clauseContext {
	var p = new(With_data_category_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_with_data_category_clause
	return p
}

func (*With_data_category_clauseContext) IsWith_data_category_clauseContext() {}

func NewWith_data_category_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *With_data_category_clauseContext {
	var p = new(With_data_category_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_with_data_category_clause

	return p
}

func (s *With_data_category_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *With_data_category_clauseContext) DATA() antlr.TerminalNode {
	return s.GetToken(SOQLParserDATA, 0)
}

func (s *With_data_category_clauseContext) CATEGORY() antlr.TerminalNode {
	return s.GetToken(SOQLParserCATEGORY, 0)
}

func (s *With_data_category_clauseContext) Data_category_spec_list() IData_category_spec_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_spec_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_category_spec_listContext)
}

func (s *With_data_category_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *With_data_category_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *With_data_category_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterWith_data_category_clause(s)
	}
}

func (s *With_data_category_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitWith_data_category_clause(s)
	}
}

func (s *With_data_category_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitWith_data_category_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) With_data_category_clause() (localctx IWith_data_category_clauseContext) {
	localctx = NewWith_data_category_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, SOQLParserRULE_with_data_category_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.Match(SOQLParserDATA)
	}
	{
		p.SetState(558)
		p.Match(SOQLParserCATEGORY)
	}
	{
		p.SetState(559)
		p.Data_category_spec_list()
	}

	return localctx
}

// IData_category_spec_listContext is an interface to support dynamic dispatch.
type IData_category_spec_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_spec_listContext differentiates from other interfaces.
	IsData_category_spec_listContext()
}

type Data_category_spec_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_spec_listContext() *Data_category_spec_listContext {
	var p = new(Data_category_spec_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_spec_list
	return p
}

func (*Data_category_spec_listContext) IsData_category_spec_listContext() {}

func NewData_category_spec_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_spec_listContext {
	var p = new(Data_category_spec_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_spec_list

	return p
}

func (s *Data_category_spec_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_spec_listContext) AllData_category_spec() []IData_category_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IData_category_specContext)(nil)).Elem())
	var tst = make([]IData_category_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IData_category_specContext)
		}
	}

	return tst
}

func (s *Data_category_spec_listContext) Data_category_spec(i int) IData_category_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IData_category_specContext)
}

func (s *Data_category_spec_listContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserAND)
}

func (s *Data_category_spec_listContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserAND, i)
}

func (s *Data_category_spec_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_spec_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_spec_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_spec_list(s)
	}
}

func (s *Data_category_spec_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_spec_list(s)
	}
}

func (s *Data_category_spec_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_spec_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_spec_list() (localctx IData_category_spec_listContext) {
	localctx = NewData_category_spec_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, SOQLParserRULE_data_category_spec_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.Data_category_spec()
	}
	p.SetState(566)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserAND {
		{
			p.SetState(562)
			p.Match(SOQLParserAND)
		}
		{
			p.SetState(563)
			p.Data_category_spec()
		}

		p.SetState(568)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IData_category_specContext is an interface to support dynamic dispatch.
type IData_category_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_specContext differentiates from other interfaces.
	IsData_category_specContext()
}

type Data_category_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_specContext() *Data_category_specContext {
	var p = new(Data_category_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_spec
	return p
}

func (*Data_category_specContext) IsData_category_specContext() {}

func NewData_category_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_specContext {
	var p = new(Data_category_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_spec

	return p
}

func (s *Data_category_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_specContext) Data_category_group_name() IData_category_group_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_group_nameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_category_group_nameContext)
}

func (s *Data_category_specContext) Data_category_selector() IData_category_selectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_selectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_category_selectorContext)
}

func (s *Data_category_specContext) Data_category_parameter_list() IData_category_parameter_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_parameter_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IData_category_parameter_listContext)
}

func (s *Data_category_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_spec(s)
	}
}

func (s *Data_category_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_spec(s)
	}
}

func (s *Data_category_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_spec() (localctx IData_category_specContext) {
	localctx = NewData_category_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, SOQLParserRULE_data_category_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(569)
		p.Data_category_group_name()
	}
	{
		p.SetState(570)
		p.Data_category_selector()
	}
	{
		p.SetState(571)
		p.Data_category_parameter_list()
	}

	return localctx
}

// IData_category_parameter_listContext is an interface to support dynamic dispatch.
type IData_category_parameter_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_parameter_listContext differentiates from other interfaces.
	IsData_category_parameter_listContext()
}

type Data_category_parameter_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_parameter_listContext() *Data_category_parameter_listContext {
	var p = new(Data_category_parameter_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_parameter_list
	return p
}

func (*Data_category_parameter_listContext) IsData_category_parameter_listContext() {}

func NewData_category_parameter_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_parameter_listContext {
	var p = new(Data_category_parameter_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_parameter_list

	return p
}

func (s *Data_category_parameter_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_parameter_listContext) AllData_category_name() []IData_category_nameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IData_category_nameContext)(nil)).Elem())
	var tst = make([]IData_category_nameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IData_category_nameContext)
		}
	}

	return tst
}

func (s *Data_category_parameter_listContext) Data_category_name(i int) IData_category_nameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IData_category_nameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IData_category_nameContext)
}

func (s *Data_category_parameter_listContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Data_category_parameter_listContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Data_category_parameter_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Data_category_parameter_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Data_category_parameter_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_parameter_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_parameter_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_parameter_list(s)
	}
}

func (s *Data_category_parameter_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_parameter_list(s)
	}
}

func (s *Data_category_parameter_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_parameter_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_parameter_list() (localctx IData_category_parameter_listContext) {
	localctx = NewData_category_parameter_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, SOQLParserRULE_data_category_parameter_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(585)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SOQLParserGROUP, SOQLParserABOVE, SOQLParserABOVE_OR_BELOW, SOQLParserAT, SOQLParserBELOW, SOQLParserCATEGORY, SOQLParserDATA, SOQLParserEND, SOQLParserOFFSET, SOQLParserORDER, SOQLParserREFERENCE, SOQLParserSCOPE, SOQLParserTRACKING, SOQLParserTHEN, SOQLParserTYPEOF, SOQLParserVIEW, SOQLParserVIEWSTAT, SOQLParserWHEN, SOQLParserCALENDAR_MONTH, SOQLParserCALENDAR_QUARTER, SOQLParserCALENDAR_YEAR, SOQLParserDAY_IN_MONTH, SOQLParserDAY_IN_WEEK, SOQLParserDAY_IN_YEAR, SOQLParserDAY_ONLY, SOQLParserFISCAL_MONTH, SOQLParserFISCAL_QUARTER, SOQLParserFISCAL_YEAR, SOQLParserHOUR_IN_DAY, SOQLParserWEEK_IN_MONTH, SOQLParserWEEK_IN_YEAR, SOQLParserAVG, SOQLParserCOUNT, SOQLParserCOUNT_DISTINCT, SOQLParserMIN, SOQLParserMAX, SOQLParserSUM, SOQLParserDISTANCE, SOQLParserGEOLOCATION, SOQLParserFORMAT, SOQLParserTOLABEL, SOQLParserCONVERT_TIME_ZONE, SOQLParserCONVERT_CURRENCY, SOQLParserGROUPING, SOQLParserYESTERDAY, SOQLParserTODAY, SOQLParserTOMORROW, SOQLParserLAST_WEEK, SOQLParserTHIS_WEEK, SOQLParserNEXT_WEEK, SOQLParserLAST_MONTH, SOQLParserTHIS_MONTH, SOQLParserNEXT_MONTH, SOQLParserLAST_90_DAYS, SOQLParserNEXT_90_DAYS, SOQLParserTHIS_QUARTER, SOQLParserLAST_QUARTER, SOQLParserNEXT_QUARTER, SOQLParserTHIS_YEAR, SOQLParserLAST_YEAR, SOQLParserNEXT_YEAR, SOQLParserTHIS_FISCAL_QUARTER, SOQLParserLAST_FISCAL_QUARTER, SOQLParserNEXT_FISCAL_QUARTER, SOQLParserTHIS_FISCAL_YEAR, SOQLParserLAST_FISCAL_YEAR, SOQLParserNEXT_FISCAL_YEAR, SOQLParserNEXT_N_DAYS, SOQLParserLAST_N_DAYS, SOQLParserN_DAYS_AGO, SOQLParserNEXT_N_WEEKS, SOQLParserLAST_N_WEEKS, SOQLParserN_WEEKS_AGO, SOQLParserNEXT_N_MONTHS, SOQLParserLAST_N_MONTHS, SOQLParserN_MONTHS_AGO, SOQLParserNEXT_N_QUARTERS, SOQLParserLAST_N_QUARTERS, SOQLParserN_QUARTERS_AGO, SOQLParserNEXT_N_YEARS, SOQLParserLAST_N_YEARS, SOQLParserN_YEARS_AGO, SOQLParserNEXT_N_FISCAL_QUARTERS, SOQLParserLAST_N_FISCAL_QUARTERS, SOQLParserN_FISCAL_QUARTERS_AGO, SOQLParserNEXT_N_FISCAL_YEARS, SOQLParserLAST_N_FISCAL_YEARS, SOQLParserN_FISCAL_YEARS_AGO, SOQLParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(573)
			p.Data_category_name()
		}

	case SOQLParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(574)
			p.Match(SOQLParserLPAREN)
		}
		{
			p.SetState(575)
			p.Data_category_name()
		}
		p.SetState(580)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SOQLParserCOMMA {
			{
				p.SetState(576)
				p.Match(SOQLParserCOMMA)
			}
			{
				p.SetState(577)
				p.Data_category_name()
			}

			p.SetState(582)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(583)
			p.Match(SOQLParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IData_category_selectorContext is an interface to support dynamic dispatch.
type IData_category_selectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsData_category_selectorContext differentiates from other interfaces.
	IsData_category_selectorContext()
}

type Data_category_selectorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyData_category_selectorContext() *Data_category_selectorContext {
	var p = new(Data_category_selectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_data_category_selector
	return p
}

func (*Data_category_selectorContext) IsData_category_selectorContext() {}

func NewData_category_selectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Data_category_selectorContext {
	var p = new(Data_category_selectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_data_category_selector

	return p
}

func (s *Data_category_selectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Data_category_selectorContext) AT() antlr.TerminalNode {
	return s.GetToken(SOQLParserAT, 0)
}

func (s *Data_category_selectorContext) ABOVE() antlr.TerminalNode {
	return s.GetToken(SOQLParserABOVE, 0)
}

func (s *Data_category_selectorContext) ABOVE_OR_BELOW() antlr.TerminalNode {
	return s.GetToken(SOQLParserABOVE_OR_BELOW, 0)
}

func (s *Data_category_selectorContext) BELOW() antlr.TerminalNode {
	return s.GetToken(SOQLParserBELOW, 0)
}

func (s *Data_category_selectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Data_category_selectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Data_category_selectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterData_category_selector(s)
	}
}

func (s *Data_category_selectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitData_category_selector(s)
	}
}

func (s *Data_category_selectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitData_category_selector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Data_category_selector() (localctx IData_category_selectorContext) {
	localctx = NewData_category_selectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, SOQLParserRULE_data_category_selector)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-34)&-(0x1f+1)) == 0 && ((1<<uint((_la-34)))&((1<<(SOQLParserABOVE-34))|(1<<(SOQLParserABOVE_OR_BELOW-34))|(1<<(SOQLParserAT-34))|(1<<(SOQLParserBELOW-34)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IGroup_by_plain_clauseContext is an interface to support dynamic dispatch.
type IGroup_by_plain_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_plain_clauseContext differentiates from other interfaces.
	IsGroup_by_plain_clauseContext()
}

type Group_by_plain_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_plain_clauseContext() *Group_by_plain_clauseContext {
	var p = new(Group_by_plain_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_group_by_plain_clause
	return p
}

func (*Group_by_plain_clauseContext) IsGroup_by_plain_clauseContext() {}

func NewGroup_by_plain_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_plain_clauseContext {
	var p = new(Group_by_plain_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_group_by_plain_clause

	return p
}

func (s *Group_by_plain_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_plain_clauseContext) Group_by_list() IGroup_by_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_listContext)
}

func (s *Group_by_plain_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_plain_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_plain_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroup_by_plain_clause(s)
	}
}

func (s *Group_by_plain_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroup_by_plain_clause(s)
	}
}

func (s *Group_by_plain_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroup_by_plain_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Group_by_plain_clause() (localctx IGroup_by_plain_clauseContext) {
	localctx = NewGroup_by_plain_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, SOQLParserRULE_group_by_plain_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		p.Group_by_list()
	}

	return localctx
}

// IGroup_by_rollup_clauseContext is an interface to support dynamic dispatch.
type IGroup_by_rollup_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_rollup_clauseContext differentiates from other interfaces.
	IsGroup_by_rollup_clauseContext()
}

type Group_by_rollup_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_rollup_clauseContext() *Group_by_rollup_clauseContext {
	var p = new(Group_by_rollup_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_group_by_rollup_clause
	return p
}

func (*Group_by_rollup_clauseContext) IsGroup_by_rollup_clauseContext() {}

func NewGroup_by_rollup_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_rollup_clauseContext {
	var p = new(Group_by_rollup_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_group_by_rollup_clause

	return p
}

func (s *Group_by_rollup_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_rollup_clauseContext) ROLLUP() antlr.TerminalNode {
	return s.GetToken(SOQLParserROLLUP, 0)
}

func (s *Group_by_rollup_clauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Group_by_rollup_clauseContext) Group_by_list() IGroup_by_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_listContext)
}

func (s *Group_by_rollup_clauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Group_by_rollup_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_rollup_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_rollup_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroup_by_rollup_clause(s)
	}
}

func (s *Group_by_rollup_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroup_by_rollup_clause(s)
	}
}

func (s *Group_by_rollup_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroup_by_rollup_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Group_by_rollup_clause() (localctx IGroup_by_rollup_clauseContext) {
	localctx = NewGroup_by_rollup_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, SOQLParserRULE_group_by_rollup_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(591)
		p.Match(SOQLParserROLLUP)
	}
	{
		p.SetState(592)
		p.Match(SOQLParserLPAREN)
	}
	{
		p.SetState(593)
		p.Group_by_list()
	}
	{
		p.SetState(594)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// IGroup_by_cube_clauseContext is an interface to support dynamic dispatch.
type IGroup_by_cube_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_cube_clauseContext differentiates from other interfaces.
	IsGroup_by_cube_clauseContext()
}

type Group_by_cube_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_cube_clauseContext() *Group_by_cube_clauseContext {
	var p = new(Group_by_cube_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_group_by_cube_clause
	return p
}

func (*Group_by_cube_clauseContext) IsGroup_by_cube_clauseContext() {}

func NewGroup_by_cube_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_cube_clauseContext {
	var p = new(Group_by_cube_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_group_by_cube_clause

	return p
}

func (s *Group_by_cube_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_cube_clauseContext) CUBE() antlr.TerminalNode {
	return s.GetToken(SOQLParserCUBE, 0)
}

func (s *Group_by_cube_clauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserLPAREN, 0)
}

func (s *Group_by_cube_clauseContext) Group_by_list() IGroup_by_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_listContext)
}

func (s *Group_by_cube_clauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SOQLParserRPAREN, 0)
}

func (s *Group_by_cube_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_cube_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_cube_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroup_by_cube_clause(s)
	}
}

func (s *Group_by_cube_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroup_by_cube_clause(s)
	}
}

func (s *Group_by_cube_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroup_by_cube_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Group_by_cube_clause() (localctx IGroup_by_cube_clauseContext) {
	localctx = NewGroup_by_cube_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 154, SOQLParserRULE_group_by_cube_clause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(596)
		p.Match(SOQLParserCUBE)
	}
	{
		p.SetState(597)
		p.Match(SOQLParserLPAREN)
	}
	{
		p.SetState(598)
		p.Group_by_list()
	}
	{
		p.SetState(599)
		p.Match(SOQLParserRPAREN)
	}

	return localctx
}

// IGroup_by_listContext is an interface to support dynamic dispatch.
type IGroup_by_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_listContext differentiates from other interfaces.
	IsGroup_by_listContext()
}

type Group_by_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_listContext() *Group_by_listContext {
	var p = new(Group_by_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_group_by_list
	return p
}

func (*Group_by_listContext) IsGroup_by_listContext() {}

func NewGroup_by_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_listContext {
	var p = new(Group_by_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_group_by_list

	return p
}

func (s *Group_by_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_listContext) AllGroup_by_spec() []IGroup_by_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IGroup_by_specContext)(nil)).Elem())
	var tst = make([]IGroup_by_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IGroup_by_specContext)
		}
	}

	return tst
}

func (s *Group_by_listContext) Group_by_spec(i int) IGroup_by_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroup_by_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IGroup_by_specContext)
}

func (s *Group_by_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Group_by_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Group_by_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroup_by_list(s)
	}
}

func (s *Group_by_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroup_by_list(s)
	}
}

func (s *Group_by_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroup_by_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Group_by_list() (localctx IGroup_by_listContext) {
	localctx = NewGroup_by_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 156, SOQLParserRULE_group_by_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		p.Group_by_spec()
	}
	p.SetState(606)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(602)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(603)
			p.Group_by_spec()
		}

		p.SetState(608)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IGroup_by_specContext is an interface to support dynamic dispatch.
type IGroup_by_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroup_by_specContext differentiates from other interfaces.
	IsGroup_by_specContext()
}

type Group_by_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroup_by_specContext() *Group_by_specContext {
	var p = new(Group_by_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_group_by_spec
	return p
}

func (*Group_by_specContext) IsGroup_by_specContext() {}

func NewGroup_by_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Group_by_specContext {
	var p = new(Group_by_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_group_by_spec

	return p
}

func (s *Group_by_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Group_by_specContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Group_by_specContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Group_by_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Group_by_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Group_by_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterGroup_by_spec(s)
	}
}

func (s *Group_by_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitGroup_by_spec(s)
	}
}

func (s *Group_by_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitGroup_by_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Group_by_spec() (localctx IGroup_by_specContext) {
	localctx = NewGroup_by_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 158, SOQLParserRULE_group_by_spec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(611)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(609)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(610)
			p.Function_call()
		}

	}

	return localctx
}

// IOrder_by_listContext is an interface to support dynamic dispatch.
type IOrder_by_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrder_by_listContext differentiates from other interfaces.
	IsOrder_by_listContext()
}

type Order_by_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_listContext() *Order_by_listContext {
	var p = new(Order_by_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_order_by_list
	return p
}

func (*Order_by_listContext) IsOrder_by_listContext() {}

func NewOrder_by_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_listContext {
	var p = new(Order_by_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_order_by_list

	return p
}

func (s *Order_by_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_listContext) AllOrder_by_spec() []IOrder_by_specContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrder_by_specContext)(nil)).Elem())
	var tst = make([]IOrder_by_specContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrder_by_specContext)
		}
	}

	return tst
}

func (s *Order_by_listContext) Order_by_spec(i int) IOrder_by_specContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrder_by_specContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrder_by_specContext)
}

func (s *Order_by_listContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SOQLParserCOMMA)
}

func (s *Order_by_listContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SOQLParserCOMMA, i)
}

func (s *Order_by_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrder_by_list(s)
	}
}

func (s *Order_by_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrder_by_list(s)
	}
}

func (s *Order_by_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrder_by_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Order_by_list() (localctx IOrder_by_listContext) {
	localctx = NewOrder_by_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 160, SOQLParserRULE_order_by_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(613)
		p.Order_by_spec()
	}
	p.SetState(618)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SOQLParserCOMMA {
		{
			p.SetState(614)
			p.Match(SOQLParserCOMMA)
		}
		{
			p.SetState(615)
			p.Order_by_spec()
		}

		p.SetState(620)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrder_by_specContext is an interface to support dynamic dispatch.
type IOrder_by_specContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrder_by_specContext differentiates from other interfaces.
	IsOrder_by_specContext()
}

type Order_by_specContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_specContext() *Order_by_specContext {
	var p = new(Order_by_specContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_order_by_spec
	return p
}

func (*Order_by_specContext) IsOrder_by_specContext() {}

func NewOrder_by_specContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_specContext {
	var p = new(Order_by_specContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_order_by_spec

	return p
}

func (s *Order_by_specContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_specContext) Order_by_field() IOrder_by_fieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrder_by_fieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrder_by_fieldContext)
}

func (s *Order_by_specContext) Order_by_direction_clause() IOrder_by_direction_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrder_by_direction_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrder_by_direction_clauseContext)
}

func (s *Order_by_specContext) Order_by_nulls_clause() IOrder_by_nulls_clauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrder_by_nulls_clauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrder_by_nulls_clauseContext)
}

func (s *Order_by_specContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_specContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_specContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrder_by_spec(s)
	}
}

func (s *Order_by_specContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrder_by_spec(s)
	}
}

func (s *Order_by_specContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrder_by_spec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Order_by_spec() (localctx IOrder_by_specContext) {
	localctx = NewOrder_by_specContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 162, SOQLParserRULE_order_by_spec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(621)
		p.Order_by_field()
	}
	p.SetState(623)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserASC || _la == SOQLParserDESC {
		{
			p.SetState(622)
			p.Order_by_direction_clause()
		}

	}
	p.SetState(626)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SOQLParserNULLS {
		{
			p.SetState(625)
			p.Order_by_nulls_clause()
		}

	}

	return localctx
}

// IOrder_by_direction_clauseContext is an interface to support dynamic dispatch.
type IOrder_by_direction_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrder_by_direction_clauseContext differentiates from other interfaces.
	IsOrder_by_direction_clauseContext()
}

type Order_by_direction_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_direction_clauseContext() *Order_by_direction_clauseContext {
	var p = new(Order_by_direction_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_order_by_direction_clause
	return p
}

func (*Order_by_direction_clauseContext) IsOrder_by_direction_clauseContext() {}

func NewOrder_by_direction_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_direction_clauseContext {
	var p = new(Order_by_direction_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_order_by_direction_clause

	return p
}

func (s *Order_by_direction_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_direction_clauseContext) ASC() antlr.TerminalNode {
	return s.GetToken(SOQLParserASC, 0)
}

func (s *Order_by_direction_clauseContext) DESC() antlr.TerminalNode {
	return s.GetToken(SOQLParserDESC, 0)
}

func (s *Order_by_direction_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_direction_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_direction_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrder_by_direction_clause(s)
	}
}

func (s *Order_by_direction_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrder_by_direction_clause(s)
	}
}

func (s *Order_by_direction_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrder_by_direction_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Order_by_direction_clause() (localctx IOrder_by_direction_clauseContext) {
	localctx = NewOrder_by_direction_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 164, SOQLParserRULE_order_by_direction_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(628)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserASC || _la == SOQLParserDESC) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrder_by_nulls_clauseContext is an interface to support dynamic dispatch.
type IOrder_by_nulls_clauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrder_by_nulls_clauseContext differentiates from other interfaces.
	IsOrder_by_nulls_clauseContext()
}

type Order_by_nulls_clauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_nulls_clauseContext() *Order_by_nulls_clauseContext {
	var p = new(Order_by_nulls_clauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_order_by_nulls_clause
	return p
}

func (*Order_by_nulls_clauseContext) IsOrder_by_nulls_clauseContext() {}

func NewOrder_by_nulls_clauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_nulls_clauseContext {
	var p = new(Order_by_nulls_clauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_order_by_nulls_clause

	return p
}

func (s *Order_by_nulls_clauseContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_nulls_clauseContext) NULLS() antlr.TerminalNode {
	return s.GetToken(SOQLParserNULLS, 0)
}

func (s *Order_by_nulls_clauseContext) FIRST() antlr.TerminalNode {
	return s.GetToken(SOQLParserFIRST, 0)
}

func (s *Order_by_nulls_clauseContext) LAST() antlr.TerminalNode {
	return s.GetToken(SOQLParserLAST, 0)
}

func (s *Order_by_nulls_clauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_nulls_clauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_nulls_clauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrder_by_nulls_clause(s)
	}
}

func (s *Order_by_nulls_clauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrder_by_nulls_clause(s)
	}
}

func (s *Order_by_nulls_clauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrder_by_nulls_clause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Order_by_nulls_clause() (localctx IOrder_by_nulls_clauseContext) {
	localctx = NewOrder_by_nulls_clauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 166, SOQLParserRULE_order_by_nulls_clause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(630)
		p.Match(SOQLParserNULLS)
	}
	{
		p.SetState(631)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserFIRST || _la == SOQLParserLAST) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOrder_by_fieldContext is an interface to support dynamic dispatch.
type IOrder_by_fieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrder_by_fieldContext differentiates from other interfaces.
	IsOrder_by_fieldContext()
}

type Order_by_fieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrder_by_fieldContext() *Order_by_fieldContext {
	var p = new(Order_by_fieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_order_by_field
	return p
}

func (*Order_by_fieldContext) IsOrder_by_fieldContext() {}

func NewOrder_by_fieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Order_by_fieldContext {
	var p = new(Order_by_fieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_order_by_field

	return p
}

func (s *Order_by_fieldContext) GetParser() antlr.Parser { return s.parser }

func (s *Order_by_fieldContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *Order_by_fieldContext) Function_call() IFunction_callContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunction_callContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunction_callContext)
}

func (s *Order_by_fieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Order_by_fieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Order_by_fieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterOrder_by_field(s)
	}
}

func (s *Order_by_fieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitOrder_by_field(s)
	}
}

func (s *Order_by_fieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitOrder_by_field(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Order_by_field() (localctx IOrder_by_fieldContext) {
	localctx = NewOrder_by_fieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 168, SOQLParserRULE_order_by_field)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(635)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(633)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(634)
			p.Function_call()
		}

	}

	return localctx
}

// IFor_valueContext is an interface to support dynamic dispatch.
type IFor_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFor_valueContext differentiates from other interfaces.
	IsFor_valueContext()
}

type For_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFor_valueContext() *For_valueContext {
	var p = new(For_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_for_value
	return p
}

func (*For_valueContext) IsFor_valueContext() {}

func NewFor_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *For_valueContext {
	var p = new(For_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_for_value

	return p
}

func (s *For_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *For_valueContext) REFERENCE() antlr.TerminalNode {
	return s.GetToken(SOQLParserREFERENCE, 0)
}

func (s *For_valueContext) VIEW() antlr.TerminalNode {
	return s.GetToken(SOQLParserVIEW, 0)
}

func (s *For_valueContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(SOQLParserUPDATE, 0)
}

func (s *For_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *For_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *For_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterFor_value(s)
	}
}

func (s *For_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitFor_value(s)
	}
}

func (s *For_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitFor_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) For_value() (localctx IFor_valueContext) {
	localctx = NewFor_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 170, SOQLParserRULE_for_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(637)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SOQLParserUPDATE-33))|(1<<(SOQLParserREFERENCE-33))|(1<<(SOQLParserVIEW-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IUpdate_valueContext is an interface to support dynamic dispatch.
type IUpdate_valueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdate_valueContext differentiates from other interfaces.
	IsUpdate_valueContext()
}

type Update_valueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdate_valueContext() *Update_valueContext {
	var p = new(Update_valueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SOQLParserRULE_update_value
	return p
}

func (*Update_valueContext) IsUpdate_valueContext() {}

func NewUpdate_valueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Update_valueContext {
	var p = new(Update_valueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SOQLParserRULE_update_value

	return p
}

func (s *Update_valueContext) GetParser() antlr.Parser { return s.parser }

func (s *Update_valueContext) TRACKING() antlr.TerminalNode {
	return s.GetToken(SOQLParserTRACKING, 0)
}

func (s *Update_valueContext) VIEWSTAT() antlr.TerminalNode {
	return s.GetToken(SOQLParserVIEWSTAT, 0)
}

func (s *Update_valueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Update_valueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Update_valueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.EnterUpdate_value(s)
	}
}

func (s *Update_valueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SOQLListener); ok {
		listenerT.ExitUpdate_value(s)
	}
}

func (s *Update_valueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SOQLVisitor:
		return t.VisitUpdate_value(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SOQLParser) Update_value() (localctx IUpdate_valueContext) {
	localctx = NewUpdate_valueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 172, SOQLParserRULE_update_value)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(639)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SOQLParserTRACKING || _la == SOQLParserVIEWSTAT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
