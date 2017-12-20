# cpgame
文字版吃鸡 for discord

## 环境设定

- 根据参与人数生成矩形的二维空间作为游戏场地。
- 空间内均匀分布AxB个房间。
- 每个房间有东南西北四个通道与相邻房间相连。
- 每经过一段时间，所有房间的位置会随机交换。
- 每经过一段时间，外围一定范围的房间会坠落，坠落前几个回合出现提示（留在其中的玩家直接死亡）
- 房间坠落至只剩最后一个房间时，若一定时间内游戏仍未结束，则强制平局
- 房间内随机分布有武器、补给和各类物品

## 游戏设定

- 游戏为多人半回合制
- 每个回合，每名玩家有固定数量的`行动点`，进行大部分操作均需使用`行动点`
- 玩家相遇时，同时开放接口进行操作，先进行操作的玩家获得先手
	- 即：第一个玩家如果发出了攻击指令，则第二个玩家就不能选择主动进攻的操作，而只能应对或选择其他操作
- 行动点消耗完后，玩家在本回合则不能进行需要消耗行动点的主动操作
- 当地图上只剩下一名玩家存活时，游戏结束，结算得分
- 每个房间分为东南西北四个局部区域，用于设定各玩家的相对位置（近身或远程）
	- 玩家不能主动在局部区域间移动，可以通过`近身`和`远离`操作相对另一玩家移动

## 数值设定

- 玩家具有相同的5点生命值
- 每回合具有5个行动点
- 肉搏消耗1个行动点，造成0.5-1.0点伤害，即两名玩家肉搏死斗决胜一般会在2回合结束
- 搜索房间消耗1个行动点，搜索房间后才能采集物资
- 采集房间内物资不消耗行动点，但属于主动操作
- 换持装备不消耗行动点，但属于主动操作
- 只有主手装备可以直接攻击，所以换持合适的装备有利于取得先手（输入命令更短）
- 行动执行需消耗时间，单次行动一般在10秒内完成
