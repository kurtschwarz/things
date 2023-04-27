import { FieldPolicy, FieldReadFunction, TypePolicies, TypePolicy } from '@apollo/client/cache';
export type AssetKeySpecifier = ('assetTags' | 'children' | 'id' | 'location' | 'locationID' | 'modelNumber' | 'name' | 'parent' | 'parentID' | 'quantity' | 'serialNumber' | 'tags' | AssetKeySpecifier)[];
export type AssetFieldPolicy = {
	assetTags?: FieldPolicy<any> | FieldReadFunction<any>,
	children?: FieldPolicy<any> | FieldReadFunction<any>,
	id?: FieldPolicy<any> | FieldReadFunction<any>,
	location?: FieldPolicy<any> | FieldReadFunction<any>,
	locationID?: FieldPolicy<any> | FieldReadFunction<any>,
	modelNumber?: FieldPolicy<any> | FieldReadFunction<any>,
	name?: FieldPolicy<any> | FieldReadFunction<any>,
	parent?: FieldPolicy<any> | FieldReadFunction<any>,
	parentID?: FieldPolicy<any> | FieldReadFunction<any>,
	quantity?: FieldPolicy<any> | FieldReadFunction<any>,
	serialNumber?: FieldPolicy<any> | FieldReadFunction<any>,
	tags?: FieldPolicy<any> | FieldReadFunction<any>
};
export type AssetConnectionKeySpecifier = ('edges' | 'pageInfo' | 'totalCount' | AssetConnectionKeySpecifier)[];
export type AssetConnectionFieldPolicy = {
	edges?: FieldPolicy<any> | FieldReadFunction<any>,
	pageInfo?: FieldPolicy<any> | FieldReadFunction<any>,
	totalCount?: FieldPolicy<any> | FieldReadFunction<any>
};
export type AssetEdgeKeySpecifier = ('cursor' | 'node' | AssetEdgeKeySpecifier)[];
export type AssetEdgeFieldPolicy = {
	cursor?: FieldPolicy<any> | FieldReadFunction<any>,
	node?: FieldPolicy<any> | FieldReadFunction<any>
};
export type AssetTagKeySpecifier = ('asset' | 'assetID' | 'id' | 'tag' | 'tagID' | AssetTagKeySpecifier)[];
export type AssetTagFieldPolicy = {
	asset?: FieldPolicy<any> | FieldReadFunction<any>,
	assetID?: FieldPolicy<any> | FieldReadFunction<any>,
	id?: FieldPolicy<any> | FieldReadFunction<any>,
	tag?: FieldPolicy<any> | FieldReadFunction<any>,
	tagID?: FieldPolicy<any> | FieldReadFunction<any>
};
export type LocationKeySpecifier = ('children' | 'deletedAt' | 'description' | 'id' | 'name' | 'parent' | 'parentID' | 'stats' | LocationKeySpecifier)[];
export type LocationFieldPolicy = {
	children?: FieldPolicy<any> | FieldReadFunction<any>,
	deletedAt?: FieldPolicy<any> | FieldReadFunction<any>,
	description?: FieldPolicy<any> | FieldReadFunction<any>,
	id?: FieldPolicy<any> | FieldReadFunction<any>,
	name?: FieldPolicy<any> | FieldReadFunction<any>,
	parent?: FieldPolicy<any> | FieldReadFunction<any>,
	parentID?: FieldPolicy<any> | FieldReadFunction<any>,
	stats?: FieldPolicy<any> | FieldReadFunction<any>
};
export type LocationConnectionKeySpecifier = ('edges' | 'pageInfo' | 'totalCount' | LocationConnectionKeySpecifier)[];
export type LocationConnectionFieldPolicy = {
	edges?: FieldPolicy<any> | FieldReadFunction<any>,
	pageInfo?: FieldPolicy<any> | FieldReadFunction<any>,
	totalCount?: FieldPolicy<any> | FieldReadFunction<any>
};
export type LocationEdgeKeySpecifier = ('cursor' | 'node' | LocationEdgeKeySpecifier)[];
export type LocationEdgeFieldPolicy = {
	cursor?: FieldPolicy<any> | FieldReadFunction<any>,
	node?: FieldPolicy<any> | FieldReadFunction<any>
};
export type LocationStatsKeySpecifier = ('totalItems' | 'totalLocations' | 'totalValue' | LocationStatsKeySpecifier)[];
export type LocationStatsFieldPolicy = {
	totalItems?: FieldPolicy<any> | FieldReadFunction<any>,
	totalLocations?: FieldPolicy<any> | FieldReadFunction<any>,
	totalValue?: FieldPolicy<any> | FieldReadFunction<any>
};
export type MutationKeySpecifier = ('createAsset' | 'createLocation' | 'createTag' | 'createUser' | 'deleteAsset' | 'deleteLocation' | 'updateAsset' | 'updateLocation' | 'updateTag' | 'updateUser' | MutationKeySpecifier)[];
export type MutationFieldPolicy = {
	createAsset?: FieldPolicy<any> | FieldReadFunction<any>,
	createLocation?: FieldPolicy<any> | FieldReadFunction<any>,
	createTag?: FieldPolicy<any> | FieldReadFunction<any>,
	createUser?: FieldPolicy<any> | FieldReadFunction<any>,
	deleteAsset?: FieldPolicy<any> | FieldReadFunction<any>,
	deleteLocation?: FieldPolicy<any> | FieldReadFunction<any>,
	updateAsset?: FieldPolicy<any> | FieldReadFunction<any>,
	updateLocation?: FieldPolicy<any> | FieldReadFunction<any>,
	updateTag?: FieldPolicy<any> | FieldReadFunction<any>,
	updateUser?: FieldPolicy<any> | FieldReadFunction<any>
};
export type NodeKeySpecifier = ('id' | NodeKeySpecifier)[];
export type NodeFieldPolicy = {
	id?: FieldPolicy<any> | FieldReadFunction<any>
};
export type PageInfoKeySpecifier = ('endCursor' | 'hasNextPage' | 'hasPreviousPage' | 'startCursor' | PageInfoKeySpecifier)[];
export type PageInfoFieldPolicy = {
	endCursor?: FieldPolicy<any> | FieldReadFunction<any>,
	hasNextPage?: FieldPolicy<any> | FieldReadFunction<any>,
	hasPreviousPage?: FieldPolicy<any> | FieldReadFunction<any>,
	startCursor?: FieldPolicy<any> | FieldReadFunction<any>
};
export type QueryKeySpecifier = ('asset' | 'assets' | 'location' | 'locations' | 'node' | 'nodes' | 'tags' | 'users' | QueryKeySpecifier)[];
export type QueryFieldPolicy = {
	asset?: FieldPolicy<any> | FieldReadFunction<any>,
	assets?: FieldPolicy<any> | FieldReadFunction<any>,
	location?: FieldPolicy<any> | FieldReadFunction<any>,
	locations?: FieldPolicy<any> | FieldReadFunction<any>,
	node?: FieldPolicy<any> | FieldReadFunction<any>,
	nodes?: FieldPolicy<any> | FieldReadFunction<any>,
	tags?: FieldPolicy<any> | FieldReadFunction<any>,
	users?: FieldPolicy<any> | FieldReadFunction<any>
};
export type TagKeySpecifier = ('asset' | 'assetTag' | 'id' | 'name' | TagKeySpecifier)[];
export type TagFieldPolicy = {
	asset?: FieldPolicy<any> | FieldReadFunction<any>,
	assetTag?: FieldPolicy<any> | FieldReadFunction<any>,
	id?: FieldPolicy<any> | FieldReadFunction<any>,
	name?: FieldPolicy<any> | FieldReadFunction<any>
};
export type TagConnectionKeySpecifier = ('edges' | 'pageInfo' | 'totalCount' | TagConnectionKeySpecifier)[];
export type TagConnectionFieldPolicy = {
	edges?: FieldPolicy<any> | FieldReadFunction<any>,
	pageInfo?: FieldPolicy<any> | FieldReadFunction<any>,
	totalCount?: FieldPolicy<any> | FieldReadFunction<any>
};
export type TagEdgeKeySpecifier = ('cursor' | 'node' | TagEdgeKeySpecifier)[];
export type TagEdgeFieldPolicy = {
	cursor?: FieldPolicy<any> | FieldReadFunction<any>,
	node?: FieldPolicy<any> | FieldReadFunction<any>
};
export type UserKeySpecifier = ('children' | 'id' | 'name' | 'parent' | 'parentID' | UserKeySpecifier)[];
export type UserFieldPolicy = {
	children?: FieldPolicy<any> | FieldReadFunction<any>,
	id?: FieldPolicy<any> | FieldReadFunction<any>,
	name?: FieldPolicy<any> | FieldReadFunction<any>,
	parent?: FieldPolicy<any> | FieldReadFunction<any>,
	parentID?: FieldPolicy<any> | FieldReadFunction<any>
};
export type UserConnectionKeySpecifier = ('edges' | 'pageInfo' | 'totalCount' | UserConnectionKeySpecifier)[];
export type UserConnectionFieldPolicy = {
	edges?: FieldPolicy<any> | FieldReadFunction<any>,
	pageInfo?: FieldPolicy<any> | FieldReadFunction<any>,
	totalCount?: FieldPolicy<any> | FieldReadFunction<any>
};
export type UserEdgeKeySpecifier = ('cursor' | 'node' | UserEdgeKeySpecifier)[];
export type UserEdgeFieldPolicy = {
	cursor?: FieldPolicy<any> | FieldReadFunction<any>,
	node?: FieldPolicy<any> | FieldReadFunction<any>
};
export type StrictTypedTypePolicies = {
	Asset?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | AssetKeySpecifier | (() => undefined | AssetKeySpecifier),
		fields?: AssetFieldPolicy,
	},
	AssetConnection?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | AssetConnectionKeySpecifier | (() => undefined | AssetConnectionKeySpecifier),
		fields?: AssetConnectionFieldPolicy,
	},
	AssetEdge?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | AssetEdgeKeySpecifier | (() => undefined | AssetEdgeKeySpecifier),
		fields?: AssetEdgeFieldPolicy,
	},
	AssetTag?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | AssetTagKeySpecifier | (() => undefined | AssetTagKeySpecifier),
		fields?: AssetTagFieldPolicy,
	},
	Location?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | LocationKeySpecifier | (() => undefined | LocationKeySpecifier),
		fields?: LocationFieldPolicy,
	},
	LocationConnection?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | LocationConnectionKeySpecifier | (() => undefined | LocationConnectionKeySpecifier),
		fields?: LocationConnectionFieldPolicy,
	},
	LocationEdge?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | LocationEdgeKeySpecifier | (() => undefined | LocationEdgeKeySpecifier),
		fields?: LocationEdgeFieldPolicy,
	},
	LocationStats?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | LocationStatsKeySpecifier | (() => undefined | LocationStatsKeySpecifier),
		fields?: LocationStatsFieldPolicy,
	},
	Mutation?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | MutationKeySpecifier | (() => undefined | MutationKeySpecifier),
		fields?: MutationFieldPolicy,
	},
	Node?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | NodeKeySpecifier | (() => undefined | NodeKeySpecifier),
		fields?: NodeFieldPolicy,
	},
	PageInfo?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | PageInfoKeySpecifier | (() => undefined | PageInfoKeySpecifier),
		fields?: PageInfoFieldPolicy,
	},
	Query?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | QueryKeySpecifier | (() => undefined | QueryKeySpecifier),
		fields?: QueryFieldPolicy,
	},
	Tag?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | TagKeySpecifier | (() => undefined | TagKeySpecifier),
		fields?: TagFieldPolicy,
	},
	TagConnection?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | TagConnectionKeySpecifier | (() => undefined | TagConnectionKeySpecifier),
		fields?: TagConnectionFieldPolicy,
	},
	TagEdge?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | TagEdgeKeySpecifier | (() => undefined | TagEdgeKeySpecifier),
		fields?: TagEdgeFieldPolicy,
	},
	User?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | UserKeySpecifier | (() => undefined | UserKeySpecifier),
		fields?: UserFieldPolicy,
	},
	UserConnection?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | UserConnectionKeySpecifier | (() => undefined | UserConnectionKeySpecifier),
		fields?: UserConnectionFieldPolicy,
	},
	UserEdge?: Omit<TypePolicy, "fields" | "keyFields"> & {
		keyFields?: false | UserEdgeKeySpecifier | (() => undefined | UserEdgeKeySpecifier),
		fields?: UserEdgeFieldPolicy,
	}
};
export type TypedTypePolicies = StrictTypedTypePolicies & TypePolicies;