/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.157
Source Server Version : 50722
Source Host           : 192.168.1.157:3306
Source Database       : sdrms

Target Server Type    : MYSQL
Target Server Version : 50722
File Encoding         : 65001

Date: 2018-08-25 17:12:32
*/

SET FOREIGN_KEY_CHECKS=0;
-- ----------------------------
-- Table structure for `kube_auth_user_cluster`
-- ----------------------------
DROP TABLE IF EXISTS `kube_auth_user_cluster`;
CREATE TABLE `kube_auth_user_cluster` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `user_type` decimal(1,0) DEFAULT NULL COMMENT '用户类型：user:0/group:1/sa:2',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='给用户授权集群查看';

-- ----------------------------
-- Records of kube_auth_user_cluster
-- ----------------------------
INSERT INTO `kube_auth_user_cluster` VALUES ('d39093fd-67c4-4cb5-959b-cf5492bf9ef5', '5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '0', '1ea98519-93b6-4854-b77d-c571af065350', '', '2018-07-11 05:46:46');
INSERT INTO `kube_auth_user_cluster` VALUES ('efe9b142-7d07-42d1-8c84-b3b4e8e2a2c5', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '0', '1', '', '2018-07-11 03:40:27');

-- ----------------------------
-- Table structure for `kube_auth_user_namespace`
-- ----------------------------
DROP TABLE IF EXISTS `kube_auth_user_namespace`;
CREATE TABLE `kube_auth_user_namespace` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL,
  `user_type` decimal(1,0) DEFAULT NULL COMMENT '用户类型：user:0/group:1/sa:2',
  `namespace_id` varchar(40) NOT NULL,
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_auth_user_namespace
-- ----------------------------
INSERT INTO `kube_auth_user_namespace` VALUES ('4987be3e-0819-406b-a623-a0f8ad7f4593', '2', '1', '93876575-86d9-43c7-af9e-d8d27a38e069', '2', '1', '2018-07-06 12:36:58');
INSERT INTO `kube_auth_user_namespace` VALUES ('4987be3e-0819-406b-a623-a0f8ad7f4594', '2', '1', '93876575-86d9-43c7-af9e-d8d27a38e069', '2', '1', '2018-07-06 12:36:58');
INSERT INTO `kube_auth_user_namespace` VALUES ('538391fc-8550-4b26-a168-17b5396584ca', '1', '0', '93876575-86d9-43c7-af9e-d8d27a38e075', '18', '1', '2018-07-10 02:40:25');
INSERT INTO `kube_auth_user_namespace` VALUES ('54806234-8daf-4aae-aa27-bb04a60fd610', '1', '0', '08ba303d-9d49-42b1-9213-1a523051bb8e', '96724a68-d10a-49f6-b852-b3c1d053e238', '1', '2018-08-17 04:38:01');
INSERT INTO `kube_auth_user_namespace` VALUES ('67f4a359-d889-4353-8101-4acbf4e473a1', '6', '0', '93876575-86d9-43c7-af9e-d8d27a38e075', '18', '1', '2018-07-10 02:40:25');
INSERT INTO `kube_auth_user_namespace` VALUES ('8d4430c1-e312-42e0-9451-b1e51937be1c', '6', '0', '93876575-86d9-43c7-af9e-d8d27a38e073', '16', '1', '2018-07-06 10:56:05');
INSERT INTO `kube_auth_user_namespace` VALUES ('9275ae2b-9a94-46ef-9d32-678fb741320d', '6', '0', '93876575-86d9-43c7-af9e-d8d27a38e075', '18', '1', '2018-07-10 02:40:25');
INSERT INTO `kube_auth_user_namespace` VALUES ('93876575-86d9-43c7-af9e-d8d27a38e070', '1', '0', '1', '1', '1', '2018-06-29 03:29:57');
INSERT INTO `kube_auth_user_namespace` VALUES ('9d05aa4e-5002-432b-b02e-e3c63a35ada5', '6', '0', '93876575-86d9-43c7-af9e-d8d27a38e075', '18', '1', '2018-07-10 02:40:25');
INSERT INTO `kube_auth_user_namespace` VALUES ('a45e7982-4323-4ea8-9ef7-7c382624659d', '1', '0', 'e236399e-14b5-4d43-9959-99b7b92ea072', '9c843224-0983-4ac3-bdc7-5395d0ec301f', '1', '2018-08-17 02:11:49');
INSERT INTO `kube_auth_user_namespace` VALUES ('abcc737f-9367-4578-ba0f-3d620a9236da', '2', '0', '93876575-86d9-43c7-af9e-d8d27a38e070', '1', '1', '2018-07-06 03:20:54');
INSERT INTO `kube_auth_user_namespace` VALUES ('b2574f22-b45c-4f9f-9c90-95574f151d9c', '1', '0', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', '1', '1', '2018-08-15 10:10:49');
INSERT INTO `kube_auth_user_namespace` VALUES ('c6992763-e50e-4340-b4ea-9f9847d04ae4', '2', '0', '93876575-86d9-43c7-af9e-d8d27a38e075', '18', '1', '2018-07-10 02:40:25');
INSERT INTO `kube_auth_user_namespace` VALUES ('d539c874-1046-4498-b61b-98077b46f97c', '6', '0', '93876575-86d9-43c7-af9e-d8d27a38e074', '17', '1', '2018-07-06 10:55:22');

-- ----------------------------
-- Table structure for `kube_bind`
-- ----------------------------
DROP TABLE IF EXISTS `kube_bind`;
CREATE TABLE `kube_bind` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(40) NOT NULL COMMENT '名称',
  `user_type` decimal(1,0) DEFAULT NULL COMMENT '用户:0/用户组:1/sa:2',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `namespace_id` varchar(40) NOT NULL COMMENT '命名空间id',
  `role_id` varchar(40) DEFAULT NULL COMMENT '角色id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户绑定角色';

-- ----------------------------
-- Records of kube_bind
-- ----------------------------
INSERT INTO `kube_bind` VALUES ('2020610a-5706-4205-b8e3-e07592d37894', '111144444', '1', '88553cf9-787f-4149-aee3-664e940f0be9', '1', '8c9821d2-d631-42ce-ba11-dc580a3740ae', '0ca29655-5663-4dae-b956-80f0864143b7');
INSERT INTO `kube_bind` VALUES ('4c6d5071-ae81-4bcd-abf3-085663d55974', '111111', '1', '3ec67537-1ed5-4243-9417-921c1dcd4f9e', '1', '93876575-86d9-43c7-af9e-d8d27a38e070', '5e8e3006-b2e7-45f9-b9c8-aaa6148299b0');
INSERT INTO `kube_bind` VALUES ('5175c6ef-e60d-4786-bc81-945332f17914', 'qqqq111', '1', '3a923331-cb1d-4b0d-bbdd-d595d404e59e', '1', '93876575-86d9-43c7-af9e-d8d27a38e070', '061ddca9-250e-4659-a8cc-d0eb08f2accf');
INSERT INTO `kube_bind` VALUES ('5844548a-518e-458a-8987-685b629e6372', 'gfd', '1', '12bf01b6-cfb5-442d-b063-993c74849bfe', '1', '8c9821d2-d631-42ce-ba11-dc580a3740ae', '0ca29655-5663-4dae-b956-80f0864143b7');
INSERT INTO `kube_bind` VALUES ('b0f806aa-f0e8-4466-85bb-037f099ad19e', 'asdf', '1', 'fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849', '2', '93876575-86d9-43c7-af9e-d8d27a38e070', '061ddca9-250e-4659-a8cc-d0eb08f2accf');

-- ----------------------------
-- Table structure for `kube_cert`
-- ----------------------------
DROP TABLE IF EXISTS `kube_cert`;
CREATE TABLE `kube_cert` (
  `id` varchar(255) NOT NULL,
  `key_name` varchar(40) NOT NULL DEFAULT '',
  `cert_value` varchar(4096) NOT NULL DEFAULT '',
  `cluster_id` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_cert
-- ----------------------------
INSERT INTO `kube_cert` VALUES ('11c4eff8-7ceb-436a-ac6e-0f00bd1d3a77', 'admin', '-----BEGIN CERTIFICATE-----\nMIIE3TCCAsWgAwIBAgIULfR+biGQGglGdOwQzNFq8xngEBQwDQYJKoZIhvcNAQEN\nBQAwZTELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0Jl\naUppbmcxDDAKBgNVBAoTA2s4czEPMA0GA1UECxMGU3lzdGVtMRMwEQYDVQQDEwpr\ndWJlcm5ldGVzMB4XDTE4MDIwNzEwNTIwMFoXDTI4MDIwNTEwNTIwMFowazELMAkG\nA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0JlaUppbmcxFzAV\nBgNVBAoTDnN5c3RlbTptYXN0ZXJzMQ8wDQYDVQQLEwZTeXN0ZW0xDjAMBgNVBAMT\nBWFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4Qxv8RKXv+5L\ncNggqutnBOdTfYWAqWmkGnmBLqyZ17q33fQWcO6GMaHrm6HkfhSBj6b9OLKBjg62\nMud8NPyRZRZ1q63H5ST0rl9qlAn63XL6lP5D6Rv9iFhK11aRNlk5eZOY1eoL6YJ5\nwXFCbMjoBOP50n8RwCCU3Q+vWWmfIVdOxSXPzTDV+p+C6z/3Xhc3teJrO36hKVM/\nLJce5BKMGellJT1Mt45hJnEk4Glw2JCjlRWujkjgiQF2XzRMv+1tMUgTDHNEjZtP\nj08VOgQeoFEfPQ9v0Cc2rh6/xvAKjoSow6h3nkJ5OfxraJlbR9poJ4JukfhMDyHC\nbX3/IzhogwIDAQABo38wfTAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFNC4PqeJLWvx\nNpV7dbx2oh00mSRkMB8GA1UdIwQYMBaAFGEf7jgOuZroWATeBfQYl2PbQczRMA0G\nCSqGSIb3DQEBDQUAA4ICAQA7spapfKPNcClwnHbNym38OXNij3lcfmzCze5zFyGR\nooOC+vbIhjzfP87BnZ80TAT4eX6/gSftARcjyFfJFPLTbnkG20Nq0DZX60yGwVku\nwS2oQQXkdMrb9z391Ih+IRezkdp7DGiwg2g6/tHntfe2fijaChDu73lJGO5uWBzr\nclmj8hk7SM+62LD4jrI+fisrh930+kxQzfRKoZaQqXnW2ZMvq0UukK5PcYtob+qw\nXMiRUNWaRKXrIUCDY5DMovT+s2YiBYEKSRdjp/AGnaDa3xm5QN2+vnimzIBbV97/\nvfx15KD8u2B2ce45nj/Wkr3VKmhieFlWg3MuqoIcsN3hoSeiu//LF8JzeGE/MRVd\nmENq/uHhoTdpxjCrnh4pPLjx0iQoX4GD9NMWBZYXeZ2p4KCgSySe7WaFyBAFoxyS\nk0U7n4jOck+j1E64Q+iVJOkW/TZgJ7VWDUldKK/fkmudf/b4oXv0/8W64eyZnbPi\n3JWqQ3O/CGG/ka+t+gyVbcQ4MuCsEeHN7o+QVicc0vUtOy/po0UDQovidutD5f0N\nk5dhd8lybNxMxrE4+2FUDJ/y56Hw8Qptc3/CsEQguq25T8qUhkGrORHSaSfaf66R\n+wwH1s/QTszaWBjqZkhT2DubhKKBMjCGgS+vdUWvob4dsqvRp/W8/KrZcYFFsT+t\n8Q==\n-----END CERTIFICATE-----\n', '1');
INSERT INTO `kube_cert` VALUES ('2bcee74c-3ef7-491e-9417-7b14e737e7ed', 'k8s-root-ca', '-----BEGIN CERTIFICATE-----\n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\n-----END CERTIFICATE-----\n', '1');
INSERT INTO `kube_cert` VALUES ('3bb5f8ae-64a0-4abe-81af-8569a66a8f9f', 'k8s-root-ca', '-----BEGIN CERTIFICATE-----\n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\n-----END CERTIFICATE-----\n', '9c843224-0983-4ac3-bdc7-5395d0ec301f');
INSERT INTO `kube_cert` VALUES ('95477b78-7411-460c-ab46-45e5b48f492d', 'admin-key', '-----BEGIN RSA PRIVATE KEY-----\nMIIEogIBAAKCAQEA4Qxv8RKXv+5LcNggqutnBOdTfYWAqWmkGnmBLqyZ17q33fQW\ncO6GMaHrm6HkfhSBj6b9OLKBjg62Mud8NPyRZRZ1q63H5ST0rl9qlAn63XL6lP5D\n6Rv9iFhK11aRNlk5eZOY1eoL6YJ5wXFCbMjoBOP50n8RwCCU3Q+vWWmfIVdOxSXP\nzTDV+p+C6z/3Xhc3teJrO36hKVM/LJce5BKMGellJT1Mt45hJnEk4Glw2JCjlRWu\njkjgiQF2XzRMv+1tMUgTDHNEjZtPj08VOgQeoFEfPQ9v0Cc2rh6/xvAKjoSow6h3\nnkJ5OfxraJlbR9poJ4JukfhMDyHCbX3/IzhogwIDAQABAoIBAGDWITTycym7N+sO\nlL/3GxUZFOvmDj7Xni1mKxgSXQBe788pwJ5HhKbGrcMCHnfCCx3tFPjYlraxzeSJ\nhwClzUpaJ7f/xwvqov6FQC8DPFcdqoWCA2bDDI0msM2ICmQAcKLHx/QECawo4ItE\nWPjGWlAAaPcShmjNnECjByjKMhb90+TGX0hAWMV2EU7xPR6J0Cbn95I2YonmMgnA\nn4GXFcgsDK3V2CpWEKqyu3/pDkwc99PYE1JylKveTTluFBwlHpMbneL+oGtBZGlq\ndw7VAZiC+bkpUkAcjiHPqdHC9RNirnVhRoPmDPXnITcF/Y3Mfo9huKGm0TWpxRhY\nzmZvJMECgYEA63hjz9Oh2oRWlCZ24e7LmgAolRytfA60DdRHqCh9SFjY+vOW0Q8S\na/o1aOxzssCccyKFuZXCpsYaH2JOULQOYfYrC45ryExozrV5o+j9rPTwRv5w7Wly\nKkqVNoNHWaJ2QDlhrxyMBHJUa0Xb2Oo66/1a8WTs4jPivqVFSbEzJtUCgYEA9Ktw\nZ/wtJfeQ/KrDTz7dYnEq7QzAiyNnF1iTpXuKmDkFxIVvhlUOH27fFQnsdoMLdz+n\nWXWsuXLYS/OSj9D2QkcmK76NFS8Y8VH22G5UE0fXfgIF5SgwTrWfXPNdrlKajmBy\nikshbHcBoG8TTzQ1agvLeV5vGbYbscUwu0MgrfcCgYBujCbP+1uRa2/6PdSyXZnY\nwxpKZxxLkduWYoMBv4CR5qR3rMSxgZH0f5Nznw7ybNsGcr61UkoAYiEBevWpjd7y\nvs+WrVaMwtKxuSFSgqAWAyiLLAl4bHjcwgcrgJaOzmcV39qsi4pwy/w2IKYGQHFJ\nObjoe6l6yUE6n/zXjAmnwQKBgCEDfQBa85Ca7hJZzE7GEcM1t/ASd2yO01tAFXQP\nzmypzRBuXNUIZwZwxGMnWqPHHOXzTdZxXWQMjgj5jb4gGQpqZUkjxg+ksj3lrGQZ\nxvhvCjGzfi4klRgZw64cHHjoJnitpObqKlFjYXHPaxCV39s2SjdPObiDbQs4q1Gp\nRiCTAoGAO5mRBihIC6btS13VFmXsKUv46xinuzWCk8rzGBQJXw69VwwC0knQZDx6\nwCCS3kmAsRVW6s2m8Y/sathNexOYR4u/l+Jj+WMAY/6Rc+kVs/ylcyOGV30t3jr0\nPrf13vqeFIJqhZp3es9a3Y975XiQm76hPCim19frD4z+YTlO6K4=\n-----END RSA PRIVATE KEY-----\n', '1');
INSERT INTO `kube_cert` VALUES ('a7b1a7ac-ded7-4fc8-805f-08ccfffa50e0', 'admin', '-----BEGIN CERTIFICATE-----\nMIIE3TCCAsWgAwIBAgIUFMywfVcd7iGGbq9DIUOhax4X3IYwDQYJKoZIhvcNAQEN\nBQAwZTELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0Jl\naUppbmcxDDAKBgNVBAoTA2s4czEPMA0GA1UECxMGU3lzdGVtMRMwEQYDVQQDEwpr\ndWJlcm5ldGVzMB4XDTE4MDEwMTA3NTMwMFoXDTI3MTIzMDA3NTMwMFowazELMAkG\nA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0JlaUppbmcxFzAV\nBgNVBAoTDnN5c3RlbTptYXN0ZXJzMQ8wDQYDVQQLEwZTeXN0ZW0xDjAMBgNVBAMT\nBWFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAulMqOB//+SMQ\nvlwQNa/ZyhRYkYCy00ugjsz20tECeq6bl/MgQAyFP8rTyZ9yZBH2WyFwhF3jhHwZ\nmRE/35PMsDoJzgTzjUWJZ/FM4StLGtq71G9CzN6Dfwa8emrAxfP/ZTevuYiJPPpu\nw+KOmfMIFtr1JC4FRW7hE22uS4MhHqcfHzwVZ0k8sVL03fe7fjznsbLmGUbjt+k2\nqLnXVfWD1zU8SiZLGUaDrIdjSvU5cZrniOb0B5j4zzyNkWdBlD2i9eh9IgZpq/KX\n64Szzk9s6g0xidQUybQk28hord5vmDjp2RaGVZsZjTH6qsp2YHGXc+MFdcUfhe4T\ngQPDdabZcQIDAQABo38wfTAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFBN1OYOkQ7FX\ng0NzskHJO7I0NZxFMB8GA1UdIwQYMBaAFJnLxofL/c7f3dR/wbpyq3T0TrX+MA0G\nCSqGSIb3DQEBDQUAA4ICAQB93t9Lk3fK9rlUPH5I3LRUz+iMXk+f000aTDxMJXMj\nszxwhPNfkOBzTYY9dCYj0PDoPbJsQaFEn9I4aYSClSOgUZeXAqGdr5UAuv61Lf50\nd7unIqSVunJHTSKaI6dpXMlDrXMi6kkhll9Z1+3jQA0OTtYANkYJNHjgqGcomVF4\nLFDmIcBDLmh3AVjjb3zOkLBldf2kn5uzdktCVC3tZ0n5Clfs8466Rl1QKFyOzVOs\nIHPV2zf6YTXDBOnpJR2UJSQ7OZ4+2kACSm0rWAjz7Jbb7A3EU+qkT/i2CD2TnNwx\nGXqYx5bWU16NgEOp/hUzeA/LM8Zt62B1ZJPmb++6jCBIq/F6HzaGRVfFyDZ7c6nD\ndFhFyBPCQPe68hCSkkOZfyACJtp07i6Ae9J7JG1DrICaFRmJ68GkJ6WdBwIszWbi\nc/VahRPVJ2pg3reSTEiggwrp+XIBEj/3ZFKT4lD0wljXehHhmqEBuRkRIQhmFy1x\nFr36Bdgbay9LIXBA66tc3DPKwlmqat/w/Lg9LBrAjbrt6wtczWgpzrE3mbs42gxo\ngK5/4Q0YT5ErPEoWHcWgx8aoySWaEmprmqtCltHrnpaGA2Un9PRUkZDmNUk4Jwaq\n3HvRLKU+57A4wVVMxXTJgRWTb5YE4evB1LAbcVmhO+dC+JSckRwEzMLDvvY6YRTb\nLA==\n-----END CERTIFICATE-----\n', '96724a68-d10a-49f6-b852-b3c1d053e238');
INSERT INTO `kube_cert` VALUES ('a9cfb386-f5de-4ba6-a3a1-9ec3b0a2f227', 'k8s-root-ca', '-----BEGIN CERTIFICATE-----\n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\n-----END CERTIFICATE-----\n', '96724a68-d10a-49f6-b852-b3c1d053e238');
INSERT INTO `kube_cert` VALUES ('e1465cc5-762a-466c-ba56-3346c805423a', 'admin', '-----BEGIN CERTIFICATE-----\nMIIE3TCCAsWgAwIBAgIUV87aQyEk4QzguEJMpYpNI2jzt/8wDQYJKoZIhvcNAQEN\nBQAwZTELMAkGA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0Jl\naUppbmcxDDAKBgNVBAoTA2s4czEPMA0GA1UECxMGU3lzdGVtMRMwEQYDVQQDEwpr\ndWJlcm5ldGVzMB4XDTE4MDczMTEwNDMwMFoXDTI4MDcyODEwNDMwMFowazELMAkG\nA1UEBhMCQ04xEDAOBgNVBAgTB0JlaUppbmcxEDAOBgNVBAcTB0JlaUppbmcxFzAV\nBgNVBAoTDnN5c3RlbTptYXN0ZXJzMQ8wDQYDVQQLEwZTeXN0ZW0xDjAMBgNVBAMT\nBWFkbWluMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA3rC+0hCLLFJ/\nsSEv6yy60AwPYFV1TOrL7n20f8RzgVRH8/QfSyZROlr828rSxlv9pN3u6+NjmjIB\n0hlS+mTnLqJNiBKihmacJRIV3fISJYvjmMwadZuhKovQoHhtFsu8gLCfMGavKimS\nV5XtRE1kgqgQh5GwbvoxhoHWAsWINNVc8cvXQnXI8yr1RO5mr8RSJzJmD6gC1dk4\ntEVeEr+HvkKWEKUgM+G2ur04d8Vfbpz5o0+75XzZarenwPZ0IhVxs/OTUZqFwCk1\nxWJ2lwmxXoo88llddFwe/r8WEf6kCeLC0xwlLbv558OwooSE0NaZRqeMLFQwPdGi\nR9+tFBbpsQIDAQABo38wfTAOBgNVHQ8BAf8EBAMCBaAwHQYDVR0lBBYwFAYIKwYB\nBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAwHQYDVR0OBBYEFOi6qZJ99A8z\nwwPytQ1Av5PaPgOsMB8GA1UdIwQYMBaAFPeVxwoUEz5TKjCMuarUzZtSH6QsMA0G\nCSqGSIb3DQEBDQUAA4ICAQBGWN9V007/uZ3e/yvqFkdEqhLvOzvOKgWPBIxZq/5N\nG3Rhip49bIOOGqwB3trm1Psk3lfAR87lfwimyHfwUzK6Gz5vIOfOaLy1KXOCs+yD\n4WNQiE+/srgjUGB8dfxZ5s/WHTsCXWR6B7RJGZcUL9+dsyYJgAn+QKhVTfKaDrxK\nFWs0n4nM8+m4WVrDJ03cLZ081dN+OzRUNzZoP1ooqOKKyEMlp9GN/bcLaf3enDZX\ngRXgfRhR0YibV1cDTAayg/27Ak5h67gL/VS6kkSlDuTw7j3yhR9z1z9ECpTa+Xe2\n+4XbOm1Yo87nglcf5l8WCPbMm0XHdF2liiipINpb910Zfsvor2jDtW+AepGOdNJy\n2qEnED3NvTtCLHM5u60vml52x0gKpJOeDan3mUrjO/Y+SnatMhqt1z0AEAzWaQQg\np0DXbnACwiEsp3p5WtGEVQawHOZaD4/2heW61nMno+mHupC94sUfd8J+4KYR63hX\nUttvd8NOAqPIbnIo/v+CQFq17HT4X7CNrNnvEKxOaffGg92iB8GGwy8eHMGRLW86\noAQM7c+rXold2frLxmpOaoXG0jK5LgNCJjKdJjVj7hjw2wNhwbBuyKGG3+LVat7e\n/Bdc65BTJa8IaZMKYaBdC5MCZ2ps8IEx3ng6zZJUVGz48pTLe6BDVLkiU3pK4OQJ\nIg==\n-----END CERTIFICATE-----\n', '9c843224-0983-4ac3-bdc7-5395d0ec301f');
INSERT INTO `kube_cert` VALUES ('f490211b-17cb-4bcd-af87-fbcfd803b1e1', 'admin-key', '-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEAulMqOB//+SMQvlwQNa/ZyhRYkYCy00ugjsz20tECeq6bl/Mg\nQAyFP8rTyZ9yZBH2WyFwhF3jhHwZmRE/35PMsDoJzgTzjUWJZ/FM4StLGtq71G9C\nzN6Dfwa8emrAxfP/ZTevuYiJPPpuw+KOmfMIFtr1JC4FRW7hE22uS4MhHqcfHzwV\nZ0k8sVL03fe7fjznsbLmGUbjt+k2qLnXVfWD1zU8SiZLGUaDrIdjSvU5cZrniOb0\nB5j4zzyNkWdBlD2i9eh9IgZpq/KX64Szzk9s6g0xidQUybQk28hord5vmDjp2RaG\nVZsZjTH6qsp2YHGXc+MFdcUfhe4TgQPDdabZcQIDAQABAoIBAQCC4qrUtD3m7+fy\nIM/ONtJhrvsscuCrpwOJbJeyLdp0/bUU/1fXRjRQRJWTm6sKese46nuUsDODvH1q\n1gAteTCfVpXOoBHKWl+UoZ/kyLnAiojJSML/EQwELmI2CEmUhLsSihSp2yy0piMq\n/To5EMzM6VCs8psHVEVNlY2LQ/j/Ayeerc0bK/gHRGs1+8KQbciDbPeThz6SnSim\nezVWflM69WkZfD60c9Mr8diCmMhPaPVy7XQWxRp/7fj/ZJHQ13DDepAxLfRBxGar\nw5mPlTlOi61k6egvWvKl6dhxkQv8+OFmu68fH3GseebGU/4ybR5MVo2gvOFO0Aih\nRPk3LiCxAoGBAMEMYGRQkmC2O4sBPl+zzEM6UPQantQLfdaMYUV6cTtGlDH3FGCC\n4L+yd9+kVa3oVMsKUFBwVTMn9CqIGs0eNfYgA7d4tiuNrf235iieWJtUdaYq7d5V\nXpYbxWN/Ksu+0DIJ9qdNl0TpIxWXXlU4fsA7jixy/efLfSWPxJFcTE0FAoGBAPcV\ng4yxnS29OFpLoh3+QMrlPlnYLsULJfFIEiZFFXK/R7e/fdDkmdPK93Bz5X/veSpQ\nRgZaoRokqe8QoU+L6mfTvtD5Ln3333ji5HlMJEGSzl498dKwCggUqQdauTEvxKVq\nka8MeLaKfgOH4eGEPQ++EuAOAKSrqKVKIF6q5KZ9AoGBALNl9AHlGlKDpwvDiRpY\ntP7Yp7hhMK5vS27hH9f4NFfYfAl+ylHx7jhW+q07Q2AMoxfYqtBFw/d5Lp+DwhxR\n8eWa1LlglTIeRM2eY2Xl2QPoVjiodksCXJb0kdenqraSyVlnBu9s3KFuYmtDMqfR\ni4DLF5FY/3m0EcWhOBO5iTKdAoGAWcZTWddfCLNrukMo0EUFPbM1iGdn9ugqDRTm\no0kOlfayFC1bhX7J9Y4VgaJajLVyDNHF36EmT91qcRZVxhVMQhVJi5w+LD7Xz5CA\n+yGTOtAgc1WGe8rCmlUHZUitaRW2GXQzIqshYRHI33eLtujZVtL1ALuVuD82s3fP\ncxYcpfECgYBgvXP+bUHmBn3wULklugXnACVoXzMRmufAUq8HgNkXCg2cYLlfRhsv\n1A6t8ej1SH0I+Y5kqTvgz/HOQijUFX1UK2cmeKl5Kic+2JZrPTOZIUnO3i+ws633\n5KEaJq785OyW01yhFAqC8lqi2QOsn8SPQI0rm8xGNhFc/eeMmD0T6Q==\n-----END RSA PRIVATE KEY-----\n', '96724a68-d10a-49f6-b852-b3c1d053e238');
INSERT INTO `kube_cert` VALUES ('fe1dcff9-4b5c-4617-b960-4a70487d39f3', 'admin-key', '-----BEGIN RSA PRIVATE KEY-----\nMIIEpAIBAAKCAQEA3rC+0hCLLFJ/sSEv6yy60AwPYFV1TOrL7n20f8RzgVRH8/Qf\nSyZROlr828rSxlv9pN3u6+NjmjIB0hlS+mTnLqJNiBKihmacJRIV3fISJYvjmMwa\ndZuhKovQoHhtFsu8gLCfMGavKimSV5XtRE1kgqgQh5GwbvoxhoHWAsWINNVc8cvX\nQnXI8yr1RO5mr8RSJzJmD6gC1dk4tEVeEr+HvkKWEKUgM+G2ur04d8Vfbpz5o0+7\n5XzZarenwPZ0IhVxs/OTUZqFwCk1xWJ2lwmxXoo88llddFwe/r8WEf6kCeLC0xwl\nLbv558OwooSE0NaZRqeMLFQwPdGiR9+tFBbpsQIDAQABAoIBAETTPXaxFEJzkPGO\nvz+hCG/Kemocn29RVgv0n/epIMXE5aQZPB6+zfYKLwJDMleKEN+GlBQlqB/8+qET\noJiw3N6F3Q6EK+T8C6BNcaKx8TfXf1U3J6pXmj0LD0S5U4XrG+xuKhUMd3DBfnBB\nbx5gCQN0q/8qYOw7uVjIAfvDTzB8eTssx3eLkHbGOR/NOxo6Xej+FiJeom9kr/vp\nQlM5f60Veq2TdhthLKo3hQGMQpPHkBs/RnZfcugIi61qvP85rx1E5ieaiW0jWrB6\ntq6Y5hLCzS/qdYA+Hb5D7bH/nFyUDoOtLE1ZfNKUZvbfWgHB3HzJ/MOfxFVaWuWo\nWNpDLgECgYEA4kcxqzP/aosYBC4TdPlFrUovW7FS9+NvGYq5x3Gn24/UbY5vWQsr\nVxnXopg3yl+8mAmm5t2GeVRknXbPtUKZIVmQH2V3grNeUNoAgPJaYMxQ6guqzepv\nCaDrSBa6A6AXueOcCyRIjwpf9985/ZewhKTcLRs7rWsFPLAeP3IrGZECgYEA+/Dp\nfEVtqRXTBAWfvrjQd1ilK41m95V3g574RFTh5UBGP2wT4nLFxGFo1k4m/bhPqvVj\nMhuq4UNJpzbSP+dM6fQj1gtLogAUpOSkG/kXpxNoQKd3gCX+2pdM5eHsXWzIgYa3\nMdrD/OsmgT/hhrDnUB1ot9WZEvFZuKr2q1cfviECgYEA2l1ULGCwsxPqKFaCxlas\njA/UZgtZAwoDtEVxBWzETZmeqd9Tyz2BJLw2oZ198Zm0OZDO9WqAlGQB+QeoaMcN\nWebBs9rKm2IXubS32biHyXRC/aomujLr1wHpLJdqCYecffKOKx5nu0qK1H0izHxv\nh4JFTG5EiBWIZ0ma1yWJPFECgYEAixXQkD5z76iQueOw8MVusLRLWuPROFVXiV41\niOOjYcA+B71OrPDXpTZxff3fIKqjsKmPfZYwm/NdseZd49F5cJ7LOds6gdCxlOZ0\ngszc9euM3kSVgDV5oItudGpo5pqrhnYspGU1VWcr9qahho4a5OuXaAWPYBnFgmE/\nlA+hAoECgYBlFr4/azi7w0DK/9gU0qIBWjfK1VDS97qSypCvlB2M9ZEwKjrxc1GR\nEcXJX4iNvbs2WRc4G84VxsgAmEQooqTid6/dHQIup3nzfPSvlj8+Aigv+Vud31AC\n2yDza0gRUMRfCIoD/dCVn7JAnf8lHBGyuG6idpZG0eJh356Y7ihCng==\n-----END RSA PRIVATE KEY-----\n', '9c843224-0983-4ac3-bdc7-5395d0ec301f');

-- ----------------------------
-- Table structure for `kube_cluster`
-- ----------------------------
DROP TABLE IF EXISTS `kube_cluster`;
CREATE TABLE `kube_cluster` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(40) NOT NULL COMMENT '名称',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k8s集群';

-- ----------------------------
-- Records of kube_cluster
-- ----------------------------
INSERT INTO `kube_cluster` VALUES ('1', 'cluster-dl11', 'daolin', '2018-06-29 09:23:39', 'for test');
INSERT INTO `kube_cluster` VALUES ('1ea98519-93b6-4854-b77d-c571af065350', 'jikytest', '', '2018-07-11 05:38:34', 'for test');
INSERT INTO `kube_cluster` VALUES ('96724a68-d10a-49f6-b852-b3c1d053e238', 'mytest', 'admin', '2018-08-17 04:30:25', '12323');
INSERT INTO `kube_cluster` VALUES ('9c843224-0983-4ac3-bdc7-5395d0ec301f', '60.100', 'admin', '2018-08-15 10:15:59', '');

-- ----------------------------
-- Table structure for `kube_env_user_cluster`
-- ----------------------------
DROP TABLE IF EXISTS `kube_env_user_cluster`;
CREATE TABLE `kube_env_user_cluster` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='环境设置-当前用户对应的集群';

-- ----------------------------
-- Records of kube_env_user_cluster
-- ----------------------------
INSERT INTO `kube_env_user_cluster` VALUES ('8111a62d-2a84-4228-8811-c99807d4826c', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '2');
INSERT INTO `kube_env_user_cluster` VALUES ('87d4f36a-a5bf-4038-bf69-4e2b460f3516', '1', '1');
INSERT INTO `kube_env_user_cluster` VALUES ('a167ca6a-58f7-478c-8914-7c52454eb80b', '5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '1ea98519-93b6-4854-b77d-c571af065350');
INSERT INTO `kube_env_user_cluster` VALUES ('c8b547cd-62a0-4a83-ac5f-8089a936ab44', '6', '2');
INSERT INTO `kube_env_user_cluster` VALUES ('dfee64b7-2577-4a2d-a23c-e323b0dac4ef', '2', '1');

-- ----------------------------
-- Table structure for `kube_env_user_namespace`
-- ----------------------------
DROP TABLE IF EXISTS `kube_env_user_namespace`;
CREATE TABLE `kube_env_user_namespace` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `namespace_id` varchar(40) NOT NULL COMMENT '命名空间id',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='环境设置-当前用户所在命名空间';

-- ----------------------------
-- Records of kube_env_user_namespace
-- ----------------------------
INSERT INTO `kube_env_user_namespace` VALUES ('6ed3627a-0585-4475-b9cb-4c147e804e53', '2', 'NamespaceId12', 'ClusterId2');
INSERT INTO `kube_env_user_namespace` VALUES ('9938c8f2-58aa-484d-b5bc-e066e09993e4', '2', 'NamespaceId11', 'ClusterId1');
INSERT INTO `kube_env_user_namespace` VALUES ('d79bd8a7-3253-41fd-a02b-4aec748f7605', '1', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', '1');
INSERT INTO `kube_env_user_namespace` VALUES ('e7659ebc-b1bd-4137-a238-8c3de5e930f2', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '93876575-86d9-43c7-af9e-d8d27a38e069', '2');
INSERT INTO `kube_env_user_namespace` VALUES ('fb8dd5a0-2277-4495-8ff2-99cecbef8610', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '93876575-86d9-43c7-af9e-d8d27a38e068', '2');

-- ----------------------------
-- Table structure for `kube_host`
-- ----------------------------
DROP TABLE IF EXISTS `kube_host`;
CREATE TABLE `kube_host` (
  `id` varchar(255) NOT NULL,
  `ip` varchar(255) NOT NULL DEFAULT '',
  `lable` varchar(255) NOT NULL DEFAULT '',
  `role` varchar(255) NOT NULL DEFAULT '',
  `host_status` tinyint(1) NOT NULL DEFAULT '0',
  `host_name` varchar(255) NOT NULL DEFAULT '',
  `pass_word` varchar(255) NOT NULL DEFAULT '',
  `cluster_id` varchar(255) NOT NULL DEFAULT '',
  `create_user` varchar(255) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `remark` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `is_deploy` tinyint(1) NOT NULL DEFAULT '0',
  `user` varchar(255) NOT NULL DEFAULT '',
  `is_install_node` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_host
-- ----------------------------
INSERT INTO `kube_host` VALUES ('001b85e7-256f-439f-a31a-d3a060059989', '192.168.60.100', '', '0', '0', '', '', '9c843224-0983-4ac3-bdc7-5395d0ec301f', '', '2018-08-15 10:16:30', '', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('03e8a11d-47f3-4f50-b62d-348d293acaa0', '192.168.60.101', '', '1', '0', '', '', '9c843224-0983-4ac3-bdc7-5395d0ec301f', '', '2018-08-15 10:16:40', '', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('0c42e96a-26e6-4619-90c9-9ebb5c2a2671', '172.16.20.11', 'role:master', '0', '0', 'master1', '', '1ea98519-93b6-4854-b77d-c571af065350', '', '2018-07-11 05:39:58', 'jikytest', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('157824f2-0d78-4bbf-9171-4296922b1f64', '192.168.1.114', 'node', '1', '0', 'mmmm', '', '1', '', '2018-07-20 06:48:26', 'ffff', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('35aeb2b5-64ff-4967-827f-f7967007e767', '172.16.20.12', 'role:node', '1', '0', 'node1', '', '1ea98519-93b6-4854-b77d-c571af065350', '', '2018-07-11 05:39:23', 'jikytest', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('35ece72f-a37e-4a55-b3f9-a25c4b195901', '172.16.20.13', 'role:node', '1', '0', 'node2', '', '1ea98519-93b6-4854-b77d-c571af065350', '', '2018-07-11 05:39:44', 'jikytest', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('362b22e9-0ca6-48b6-ba17-63429631ea90', '192.168.1.153', 'master', '0', '0', 'master1', '', '1', '', '2018-07-31 02:18:28', '', '0', '0', '', '1');
INSERT INTO `kube_host` VALUES ('5374545d-1542-44eb-88d3-373c5585d9cb', '192.168.1.111', 'la la la ', '0', '0', 'test1', '', '1', '', '2018-07-10 06:53:54', '', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('9c28c8a6-7954-40d2-bac7-66f25779660e', '192.168.1.112', 'la la la 1', '0', '0', 'test2', '', '', '', '2018-07-10 06:54:04', '', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('bb409a5e-b81a-4a0b-8a37-52d210348148', '192.168.10.4', 'role:master', '0', '0', 'master1', '', '96724a68-d10a-49f6-b852-b3c1d053e238', '', '2018-08-17 04:31:02', 'dd', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('e7abe3df-4c97-433c-93b6-3bb6b119fc49', '192.168.1.113', 'la la la 2', '1', '0', 'test3', '', '', '', '2018-07-10 06:54:19', '', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('ec7cc74b-f9e3-428e-b2a3-0f6a283c10a8', '192.168.60.35', 'master', '0', '0', 'master', '', '', '', '2018-07-30 08:14:42', 'sss', '0', '0', '', '0');
INSERT INTO `kube_host` VALUES ('f868ce6a-fc06-4a9a-8554-92cd2213be68', '192.168.10.5', 'role:node', '1', '0', 'node1', '', '', '', '2018-08-17 09:49:58', 'dd', '0', '0', '', '1');

-- ----------------------------
-- Table structure for `kube_image`
-- ----------------------------
DROP TABLE IF EXISTS `kube_image`;
CREATE TABLE `kube_image` (
  `id` varchar(255) NOT NULL,
  `name` varchar(100) NOT NULL DEFAULT '',
  `tag` varchar(40) NOT NULL DEFAULT '',
  `env` varchar(100) NOT NULL DEFAULT '',
  `heartbeat` varchar(100) NOT NULL DEFAULT '',
  `runcmd` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_image
-- ----------------------------
INSERT INTO `kube_image` VALUES ('1', 'redis', 'v1', '', '', '');
INSERT INTO `kube_image` VALUES ('1c023302-6523-4f36-b007-1ad49bb7d343', 'harbor.zxbike.cn/public/tomcat7069_jdk80151', 'v1.0.1', '', '', '');
INSERT INTO `kube_image` VALUES ('2', 'redis', 'v2', '', '', '');
INSERT INTO `kube_image` VALUES ('3', 'redis-dl', 'v3', 'tail=55', 'ddsadsad', 'tai;oee');
INSERT INTO `kube_image` VALUES ('4', 'redis', 'latest', 'd', 'd', 'd');
INSERT INTO `kube_image` VALUES ('5', 'harbor.zxbike.cn/public/redis', '3.2.11', '', '', '');
INSERT INTO `kube_image` VALUES ('58dc2d2d-2e6e-4bd6-a6a0-c56237dae89a', 'codis', 'v12', 'tttttttttt', 'ttttttttttt', 'tttttttttttt');

-- ----------------------------
-- Table structure for `kube_module_instance`
-- ----------------------------
DROP TABLE IF EXISTS `kube_module_instance`;
CREATE TABLE `kube_module_instance` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(40) DEFAULT NULL COMMENT '组件名称',
  `type` decimal(2,0) DEFAULT NULL COMMENT '组件类型',
  `port` decimal(4,0) DEFAULT NULL COMMENT '端口',
  `state` decimal(1,0) DEFAULT NULL COMMENT '状态',
  `cluster_id` varchar(40) DEFAULT NULL COMMENT '集群id',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='组件实例状态';

-- ----------------------------
-- Records of kube_module_instance
-- ----------------------------

-- ----------------------------
-- Table structure for `kube_namespace`
-- ----------------------------
DROP TABLE IF EXISTS `kube_namespace`;
CREATE TABLE `kube_namespace` (
  `id` varchar(40) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `lable` varchar(100) NOT NULL DEFAULT '',
  `cluster_id` varchar(40) NOT NULL DEFAULT '',
  `create_user` varchar(40) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `remark` varchar(200) NOT NULL DEFAULT '',
  `stype` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_namespace
-- ----------------------------
INSERT INTO `kube_namespace` VALUES ('08ba303d-9d49-42b1-9213-1a523051bb8e', 'default', '', '96724a68-d10a-49f6-b852-b3c1d053e238', '1', '2018-08-17 04:37:49', '', '1');
INSERT INTO `kube_namespace` VALUES ('1c30a63a-8606-483b-8f1f-9998cdfff548', 'qiqi-zxbk-china-middleware', '', '1', '1', '2018-08-15 09:05:19', 'qiqi中间件', '1');
INSERT INTO `kube_namespace` VALUES ('27b48203-ab83-4b50-b610-e4722db3e8dd', 'JxmTest_Name', 'JxmTest_Lable', '1', 'User1', '2018-07-04 10:03:45', '测试', '1');
INSERT INTO `kube_namespace` VALUES ('286798f8-186c-4247-8633-84786149c392', 'istio-test', '', '1', '1', '2018-08-15 09:04:38', 'istio微服务测试', '1');
INSERT INTO `kube_namespace` VALUES ('2c86e42d-e16b-4929-a833-4fe7a4620be9', 'istio-system', '', '1', '1', '2018-08-15 09:04:22', 'istio-system测试', '1');
INSERT INTO `kube_namespace` VALUES ('359774b8-4921-4347-9a19-61e30933c779', 'dltest', '', '2', '1', '2018-07-16 02:20:42', '', '1');
INSERT INTO `kube_namespace` VALUES ('38a0e424-9463-4001-a1bc-d421e71a05fb', 'kube-system', '', '1', '1', '2018-08-03 09:54:25', 'xit', '1');
INSERT INTO `kube_namespace` VALUES ('93876575-86d9-43c7-af9e-d8d27a38e068', 'namespace2', 'JxmTest_Lable', '2', 'User1', '2018-06-28 07:17:30', '测试', '1');
INSERT INTO `kube_namespace` VALUES ('93876575-86d9-43c7-af9e-d8d27a38e069', 'namespace3', 'JxmTest_Lable', '2', 'User2', '2018-06-29 07:17:30', '测试', '1');
INSERT INTO `kube_namespace` VALUES ('a3f7fce9-e90f-486d-a3e6-4c38ff3f8e59', 'kubeapps', '', '1', '1', '2018-08-23 03:00:32', '153测试环境', '1');
INSERT INTO `kube_namespace` VALUES ('b0eaf648-1935-4876-9b65-60ecb1a5c562', 'default', '', '1', '1', '2018-08-02 09:52:42', '默认命名空间', '1');
INSERT INTO `kube_namespace` VALUES ('e236399e-14b5-4d43-9959-99b7b92ea072', 'default', '', '9c843224-0983-4ac3-bdc7-5395d0ec301f', '1', '2018-08-17 02:10:39', '', '1');
INSERT INTO `kube_namespace` VALUES ('f08636c8-4b04-4f53-936d-8e282063f88b', 'middleware', '', '1', '1', '2018-08-15 09:04:55', '中间件', '1');

-- ----------------------------
-- Table structure for `kube_oper_definition`
-- ----------------------------
DROP TABLE IF EXISTS `kube_oper_definition`;
CREATE TABLE `kube_oper_definition` (
  `id` varchar(255) NOT NULL,
  `oper_type` int(11) NOT NULL DEFAULT '0',
  `oper_name` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_oper_definition
-- ----------------------------
INSERT INTO `kube_oper_definition` VALUES ('1', '1', 'read');
INSERT INTO `kube_oper_definition` VALUES ('2', '2', 'write');
INSERT INTO `kube_oper_definition` VALUES ('3', '3', 'get,list');

-- ----------------------------
-- Table structure for `kube_publish_service`
-- ----------------------------
DROP TABLE IF EXISTS `kube_publish_service`;
CREATE TABLE `kube_publish_service` (
  `id` varchar(40) NOT NULL,
  `domain_name` varchar(40) NOT NULL DEFAULT '',
  `name` varchar(40) NOT NULL DEFAULT '',
  `ramark` varchar(100) NOT NULL DEFAULT '',
  `create_user` varchar(40) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `cluster_id` varchar(40) NOT NULL DEFAULT '',
  `namespace_id` varchar(40) NOT NULL DEFAULT '',
  `stype` varchar(8) NOT NULL DEFAULT '',
  `service_id` varchar(40) NOT NULL DEFAULT '',
  `deploy_name` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_publish_service
-- ----------------------------
INSERT INTO `kube_publish_service` VALUES ('2f8a01b3-e7a3-4cd1-937e-2a8afa2a4d63', '55.com', 'sasas', 'asdsa', '1', '2018-07-19 04:52:40', '1', '27b48203-ab83-4b50-b610-e4722db3e8dd', 'HTTP', '', '');
INSERT INTO `kube_publish_service` VALUES ('62a61b8d-443b-40a6-86a3-7995d7915e34', 'www', 'www', 'beiz ', '1', '2018-07-26 06:51:07', '1', '66f618a0-28ff-462e-8e12-bb85355fa9af', 'HTTP', '', '');
INSERT INTO `kube_publish_service` VALUES ('6bc1c32c-377b-4593-a815-349573d563a9', 'jxmtest.com', 'jxmtest.com', '3334423', '1', '2018-08-24 08:31:59', '1', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', 'HTTP', '', '');
INSERT INTO `kube_publish_service` VALUES ('770ce2c7-b746-4c46-9103-15fd1219e4a6', '', 'bikeauto', '', '1', '2018-08-13 09:25:08', '1', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', 'TCP', '7d8c7aa4-2b3e-4927-8d5c-1bef659f6a6f', 'bikeauto');
INSERT INTO `kube_publish_service` VALUES ('de787335-9f8d-422b-a2ef-4110b8a67c7b', '', 'sdsadsa', 'sdad', '1', '2018-07-19 04:58:54', '1', '93876575-86d9-43c7-af9e-d8d27a38e070', 'TCP', '659ed40f-aac1-4c8a-856e-9766f606e5fa', 'sdsadsa');
INSERT INTO `kube_publish_service` VALUES ('efd79377-8ca2-4669-b9c2-031d68882814', 'test.com', 'test', 'sdfsf', '1', '2018-07-19 04:52:07', '1', '66f618a0-28ff-462e-8e12-bb85355fa9af', 'HTTP', '', '');

-- ----------------------------
-- Table structure for `kube_publish_service_path`
-- ----------------------------
DROP TABLE IF EXISTS `kube_publish_service_path`;
CREATE TABLE `kube_publish_service_path` (
  `id` varchar(40) NOT NULL,
  `path` varchar(40) NOT NULL DEFAULT '',
  `service_id` varchar(40) NOT NULL DEFAULT '',
  `port_id` varchar(40) NOT NULL DEFAULT '',
  `pservice_id` varchar(40) NOT NULL DEFAULT '',
  `host_port` varchar(8) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_publish_service_path
-- ----------------------------
INSERT INTO `kube_publish_service_path` VALUES ('06f8476e-e62b-48ee-afa4-1a649a9728e1', '/paht2', '659ed40f-aac1-4c8a-856e-9766f606e5fa', '602b4107-3d5a-40b7-aa6a-2b96235ec5b2', '62a61b8d-443b-40a6-86a3-7995d7915e34', '', '2018-07-26 06:51:07');
INSERT INTO `kube_publish_service_path` VALUES ('0b520aba-b2f3-4692-b0c3-ba0ad19b802e', 'ffdf', '659ed40f-aac1-4c8a-856e-9766f606e5fa', '8a9807f7-fc6d-4e19-b8b7-852e6cbc52da', 'efd79377-8ca2-4669-b9c2-031d68882814', '', '2018-07-19 04:52:07');
INSERT INTO `kube_publish_service_path` VALUES ('15525a59-18c9-4c5d-a424-b7e342662287', '/path2', 'ed8e3583-ba62-439d-b073-442020a25f2c', '330d8cf2-3206-4e8e-9bd4-ea1cb70e099a', '6bc1c32c-377b-4593-a815-349573d563a9', '', '2018-08-24 08:31:59');
INSERT INTO `kube_publish_service_path` VALUES ('2c954119-cfa7-43fc-acb6-4b617ef53a90', '', '', '24e64361-a659-4bf8-a813-c48ab2bf5bcd', '770ce2c7-b746-4c46-9103-15fd1219e4a6', '33401', '2018-08-05 09:48:30');
INSERT INTO `kube_publish_service_path` VALUES ('32701043-3599-44e8-a028-7290322e4a75', '', '', '602b4107-3d5a-40b7-aa6a-2b96235ec5b2', 'de787335-9f8d-422b-a2ef-4110b8a67c7b', '333', '2018-07-19 04:58:54');
INSERT INTO `kube_publish_service_path` VALUES ('351fa018-2e63-4e6f-afe5-bc92913f82cb', 'sads', 'f60e2ac8-ad54-47b0-8bb2-bac0369fe792', 'ab85981c-aa91-4d67-9ed1-4627727d5091', '2f8a01b3-e7a3-4cd1-937e-2a8afa2a4d63', '', '2018-07-19 04:52:40');
INSERT INTO `kube_publish_service_path` VALUES ('8f02667d-b415-4a7e-bff0-8dd4d02ae54e', '', '', 'd1dfa3e6-8bfd-4ccf-a406-661601ade8a3', 'e202fd3b-1d48-47c1-b1d6-692f5a558bc6', '', '2018-08-09 08:56:10');
INSERT INTO `kube_publish_service_path` VALUES ('9eb8f51b-8a4d-495d-9315-af4563d7a735', 'dasd', '659ed40f-aac1-4c8a-856e-9766f606e5fa', '602b4107-3d5a-40b7-aa6a-2b96235ec5b2', 'efd79377-8ca2-4669-b9c2-031d68882814', '', '2018-07-19 04:52:07');
INSERT INTO `kube_publish_service_path` VALUES ('aae2cd3d-0fb7-40e0-befb-bcebefd68c1f', '', '', '4500f7b4-0dfc-4842-bcc0-de08d5ed7132', '770ce2c7-b746-4c46-9103-15fd1219e4a6', '33303', '2018-08-05 09:48:30');
INSERT INTO `kube_publish_service_path` VALUES ('ac6406a1-4015-44f2-926e-78cebcb05c92', '', '', '8a9807f7-fc6d-4e19-b8b7-852e6cbc52da', 'de787335-9f8d-422b-a2ef-4110b8a67c7b', '55', '2018-07-18 20:52:59');
INSERT INTO `kube_publish_service_path` VALUES ('b2542ed8-bd9b-40f6-a29b-18838ff601b5', '/path1', 'ed8e3583-ba62-439d-b073-442020a25f2c', 'fab60be7-6f8f-45de-9791-972ae0b7a840', '6bc1c32c-377b-4593-a815-349573d563a9', '', '2018-07-20 08:33:37');
INSERT INTO `kube_publish_service_path` VALUES ('b8723d32-8b9b-4a90-8108-b44424c72eb8', 'sadsd', 'f60e2ac8-ad54-47b0-8bb2-bac0369fe792', '43235744-f14c-476e-8014-7fed38a2a574', '2f8a01b3-e7a3-4cd1-937e-2a8afa2a4d63', '', '2018-07-19 04:52:40');
INSERT INTO `kube_publish_service_path` VALUES ('c6362a35-7ddc-4255-9c90-3826e78f1be8', '/path', '16f4735f-1750-4fd4-879a-a972e0ef1a3f', 'ee664bf4-e50c-468b-a331-524859c12519', '62a61b8d-443b-40a6-86a3-7995d7915e34', '', '2018-07-25 22:50:43');

-- ----------------------------
-- Table structure for `kube_resource`
-- ----------------------------
DROP TABLE IF EXISTS `kube_resource`;
CREATE TABLE `kube_resource` (
  `id` varchar(255) NOT NULL,
  `res_type` int(11) NOT NULL DEFAULT '0',
  `oper_role` varchar(100) NOT NULL DEFAULT '',
  `api_groups` varchar(1) NOT NULL DEFAULT '',
  `resource_names` varchar(40) NOT NULL DEFAULT '',
  `role_type` int(11) NOT NULL DEFAULT '0',
  `role_id` varchar(40) NOT NULL DEFAULT '',
  `user_oper_type` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_resource
-- ----------------------------
INSERT INTO `kube_resource` VALUES ('4620f318-0312-469a-8331-f81bed960f5e', '0', '', '', 'deployment', '0', '5e8e3006-b2e7-45f9-b9c8-aaa6148299b0', '2');
INSERT INTO `kube_resource` VALUES ('491ca811-d1b0-4bff-a269-4599daf1588f', '0', '', '', 'deployment', '0', '2e2e1f74-0fab-45bb-91b6-e79579b4fb5b', '3');
INSERT INTO `kube_resource` VALUES ('4edd505d-732d-4efb-a662-54bceb20f8b1', '0', '', '', 'sdf', '0', '5e8e3006-b2e7-45f9-b9c8-aaa6148299b0', '1');

-- ----------------------------
-- Table structure for `kube_role`
-- ----------------------------
DROP TABLE IF EXISTS `kube_role`;
CREATE TABLE `kube_role` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `name` varchar(40) NOT NULL COMMENT '角色名称',
  `role_type` decimal(1,0) NOT NULL COMMENT 'role:0/clusterorle:1`',
  `namespace_id` varchar(40) NOT NULL COMMENT '命名空间id',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k8s角色';

-- ----------------------------
-- Records of kube_role
-- ----------------------------
INSERT INTO `kube_role` VALUES ('09b3e701-5cba-4bdb-b9d5-7064e477cd25', '12212', '0', '', '2', '', '2018-07-04 07:50:39', '');
INSERT INTO `kube_role` VALUES ('2e2e1f74-0fab-45bb-91b6-e79579b4fb5b', 'dltest', '0', '', '1', '', '2018-07-05 02:28:53', '');
INSERT INTO `kube_role` VALUES ('3c0822eb-d22e-462f-bb8e-fe38aa083891', 'asdfasdf', '0', '', '1', 'admin', '2018-08-22 08:54:10', '');
INSERT INTO `kube_role` VALUES ('5e8e3006-b2e7-45f9-b9c8-aaa6148299b0', '123456', '0', '', '1', '', '2018-07-04 07:45:22', '');

-- ----------------------------
-- Table structure for `kube_service`
-- ----------------------------
DROP TABLE IF EXISTS `kube_service`;
CREATE TABLE `kube_service` (
  `id` varchar(255) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `image_name` varchar(40) NOT NULL DEFAULT '',
  `env` varchar(100) NOT NULL DEFAULT '',
  `run` varchar(100) NOT NULL DEFAULT '',
  `host_ip` tinyint(1) NOT NULL DEFAULT '0',
  `cpu_need` varchar(255) NOT NULL DEFAULT '',
  `cpu_max` varchar(255) NOT NULL DEFAULT '',
  `memory_need` varchar(255) NOT NULL DEFAULT '',
  `memory_max` varchar(255) NOT NULL DEFAULT '',
  `service_num` int(11) NOT NULL DEFAULT '0',
  `heartbeat` varchar(100) NOT NULL DEFAULT '',
  `run_time` int(11) NOT NULL DEFAULT '0',
  `soket_time` int(11) NOT NULL DEFAULT '0',
  `create_user` varchar(40) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `remark` varchar(200) NOT NULL DEFAULT '',
  `group_id` varchar(40) NOT NULL DEFAULT '',
  `namespace_id` varchar(40) NOT NULL DEFAULT '',
  `cluster_id` varchar(40) NOT NULL DEFAULT '',
  `is_version` int(11) NOT NULL DEFAULT '0',
  `father_id` varchar(40) NOT NULL DEFAULT '',
  `service_mark` varchar(40) NOT NULL DEFAULT '',
  `version_id` varchar(255) NOT NULL DEFAULT '',
  `image_id` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_service
-- ----------------------------
INSERT INTO `kube_service` VALUES ('7751bf4a-9190-4dfe-ad60-8de6ebc67f89', 'redis-test', 'harbor.zxbike.cn/public/redis:3.2.11', '', '', '1', '0.1', '0.2', '50', '100', '2', '', '0', '0', 'admin', '2018-08-24 09:08:30', '', '', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', '1', '0', '', '', '', '');
INSERT INTO `kube_service` VALUES ('985faa38-bd38-4e5d-a7ec-1941b51d765b', 'redis-test', '192.168.1.158/public/redis:3.2.11', '', '', '1', '0.1', '0.2', '50', '100', '2', '', '0', '0', 'admin', '2018-08-24 00:18:09', '', '', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', '1', '1', '7751bf4a-9190-4dfe-ad60-8de6ebc67f89', '', '4af65a9a-a7d8-4ff1-846d-d0a36600196c', '');
INSERT INTO `kube_service` VALUES ('ed8e3583-ba62-439d-b073-442020a25f2c', 'jxmtest', 'harbor.zxbike.cn/public/busybox:v1.0.0', '', '', '1', '0.1', '0.2', '50', '100', '1', '', '0', '0', 'admin', '2018-08-24 10:14:16', '', '', 'b0eaf648-1935-4876-9b65-60ecb1a5c562', '1', '0', '', '', '', '1c023302-6523-4f36-b007-1ad49bb7d343');

-- ----------------------------
-- Table structure for `kube_service_accounts`
-- ----------------------------
DROP TABLE IF EXISTS `kube_service_accounts`;
CREATE TABLE `kube_service_accounts` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `namespace_id` varchar(40) DEFAULT NULL COMMENT '命名空间',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  `name` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k8s  Service Accounts';

-- ----------------------------
-- Records of kube_service_accounts
-- ----------------------------
INSERT INTO `kube_service_accounts` VALUES ('3ec67537-1ed5-4243-9417-921c1dcd4f9e', '1', '', '1', '2018-07-09 02:45:40', '备注1', 'sa1');
INSERT INTO `kube_service_accounts` VALUES ('88553cf9-787f-4149-aee3-664e940f0be9', '1', '', '1', '2018-07-11 02:24:32', '111ddd', 'satest111');

-- ----------------------------
-- Table structure for `kube_service_issue`
-- ----------------------------
DROP TABLE IF EXISTS `kube_service_issue`;
CREATE TABLE `kube_service_issue` (
  `id` varchar(255) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `service_id` varchar(40) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  `remark` varchar(255) NOT NULL DEFAULT '',
  `service_mark` varchar(40) NOT NULL DEFAULT '',
  `type` varchar(6) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_service_issue
-- ----------------------------

-- ----------------------------
-- Table structure for `kube_service_port`
-- ----------------------------
DROP TABLE IF EXISTS `kube_service_port`;
CREATE TABLE `kube_service_port` (
  `id` varchar(255) NOT NULL,
  `name` varchar(40) NOT NULL DEFAULT '',
  `protocol` varchar(8) NOT NULL DEFAULT '',
  `container_port` varchar(255) NOT NULL DEFAULT '',
  `service_port` varchar(255) NOT NULL DEFAULT '',
  `is_main` varchar(255) NOT NULL DEFAULT '',
  `service_id` varchar(40) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_service_port
-- ----------------------------
INSERT INTO `kube_service_port` VALUES ('24e64361-a659-4bf8-a813-c48ab2bf5bcd', 'noweb', 'TCP', '10401', '889', 'No', '7d8c7aa4-2b3e-4927-8d5c-1bef659f6a6f');
INSERT INTO `kube_service_port` VALUES ('330d8cf2-3206-4e8e-9bd4-ea1cb70e099a', 'contraionsvc', 'TCP', '8000', '800', 'No', 'ed8e3583-ba62-439d-b073-442020a25f2c');
INSERT INTO `kube_service_port` VALUES ('402f4d5a-3c34-427f-a869-c11d21e548c4', 'MyTcp', 'TCP', '10401', '889', 'No', 'c7d13df6-5683-4c00-92c8-261f1d2ec885');
INSERT INTO `kube_service_port` VALUES ('43235744-f14c-476e-8014-7fed38a2a574', '333', 'TCP/UDP', '333', '333', 'No', 'f60e2ac8-ad54-47b0-8bb2-bac0369fe792');
INSERT INTO `kube_service_port` VALUES ('4500f7b4-0dfc-4842-bcc0-de08d5ed7132', 'web', 'TCP', '8080', '880', 'No', '7d8c7aa4-2b3e-4927-8d5c-1bef659f6a6f');
INSERT INTO `kube_service_port` VALUES ('602b4107-3d5a-40b7-aa6a-2b96235ec5b2', 'TestName', 'TCP', 'ooppo', 'qw1qw', 'Yes', '659ed40f-aac1-4c8a-856e-9766f606e5fa');
INSERT INTO `kube_service_port` VALUES ('6f23e962-0d6f-4dc0-8ce8-9060be73e7fa', 'web', 'TCP', '8080', '880', 'No', 'c7d13df6-5683-4c00-92c8-261f1d2ec885');
INSERT INTO `kube_service_port` VALUES ('8a9807f7-fc6d-4e19-b8b7-852e6cbc52da', 'serviceport', 'TCP/UDP', '3306', '3306', 'No', '659ed40f-aac1-4c8a-856e-9766f606e5fa');
INSERT INTO `kube_service_port` VALUES ('95449143-c0e3-4da5-adfb-2723815af265', '33', 'TCP', '3312', '3312', 'Yes', '41e1ff22-9f1c-483d-ae64-f5f264e5b482');
INSERT INTO `kube_service_port` VALUES ('ab85981c-aa91-4d67-9ed1-4627727d5091', '321', 'TCP', '321', '321', 'Yes', 'f60e2ac8-ad54-47b0-8bb2-bac0369fe792');
INSERT INTO `kube_service_port` VALUES ('b711ac73-f85c-43f3-a2de-4727500f0291', '22', 'TCP/UDP', '3309', '3389', 'No', '41e1ff22-9f1c-483d-ae64-f5f264e5b482');
INSERT INTO `kube_service_port` VALUES ('c052b99a-6bdb-402e-bf4e-218c9483ac18', '3213', 'TCP/UDP', '323', '3333', 'No', '1cf5d479-ada7-4134-b116-e05bbf1877af');
INSERT INTO `kube_service_port` VALUES ('d1dfa3e6-8bfd-4ccf-a406-661601ade8a3', '1212', 'TCP/UDP', '', '', 'No', '5828d1d1-244d-41b2-ad7f-89cb0ef431da');
INSERT INTO `kube_service_port` VALUES ('d3924297-6add-451a-8e26-605ea2975f34', '', 'TCP', '8080', '880', 'Yes', '5828d1d1-244d-41b2-ad7f-89cb0ef431da');
INSERT INTO `kube_service_port` VALUES ('f4d54bf1-4ce8-4774-af7f-0a8253054ea9', 'bigdata', 'TCP/UDP', '8807', '8807', 'No', '659ed40f-aac1-4c8a-856e-9766f606e5fa');
INSERT INTO `kube_service_port` VALUES ('f8b8e872-53dd-4e4e-b84a-cf554703dfe7', '', 'UDP', '8090', '890', 'No', '5828d1d1-244d-41b2-ad7f-89cb0ef431da');
INSERT INTO `kube_service_port` VALUES ('fab60be7-6f8f-45de-9791-972ae0b7a840', 'web', 'TCP', '8080', '889', 'Yes', 'ed8e3583-ba62-439d-b073-442020a25f2c');

-- ----------------------------
-- Table structure for `kube_service_version`
-- ----------------------------
DROP TABLE IF EXISTS `kube_service_version`;
CREATE TABLE `kube_service_version` (
  `id` varchar(255) NOT NULL,
  `version_name` varchar(40) NOT NULL DEFAULT '',
  `remark` varchar(255) NOT NULL DEFAULT '',
  `create_user` varchar(40) NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of kube_service_version
-- ----------------------------
INSERT INTO `kube_service_version` VALUES ('03362620-a619-4fed-8267-ffaadf5a748e', 'v3', '32132', 'admin', '2018-07-27 09:51:47');
INSERT INTO `kube_service_version` VALUES ('124bba58-33fa-44da-a67e-9f03f57e0b07', 'JxmTestAutoV1.0', '', 'admin', '2018-08-10 10:29:33');
INSERT INTO `kube_service_version` VALUES ('255c7b27-2a9f-4f82-830c-968d84ba4f88', 'v1', 'ddd', 'admin', '2018-07-26 06:47:04');
INSERT INTO `kube_service_version` VALUES ('2d05d108-05b2-41d4-9a0c-2596a37945a4', 'V1.02', '', 'admin', '2018-08-10 10:31:48');
INSERT INTO `kube_service_version` VALUES ('35f43f9f-912a-47a6-a355-a85cc55b9d96', 'v2', '', 'admin', '2018-08-10 09:03:36');
INSERT INTO `kube_service_version` VALUES ('4af65a9a-a7d8-4ff1-846d-d0a36600196c', 'v1', 'ttt', 'admin', '2018-08-24 08:21:24');
INSERT INTO `kube_service_version` VALUES ('766d28e7-7f4a-4eda-b2c6-5b9473e81498', '1.0', 'ttt', 'admin', '2018-08-10 01:13:25');
INSERT INTO `kube_service_version` VALUES ('7aa724d5-af32-4687-b043-5ac63acb0801', 'v1', '', 'admin', '2018-08-17 02:10:16');
INSERT INTO `kube_service_version` VALUES ('8b1139ae-a101-4a36-aaf7-92ea7aa3b136', 'v1.1', '', 'admin', '2018-08-10 10:35:13');
INSERT INTO `kube_service_version` VALUES ('a038a499-7747-4719-b567-93d3184f3724', 'v3.0', '', 'admin', '2018-08-10 09:12:25');
INSERT INTO `kube_service_version` VALUES ('aa22a28a-d160-44ee-9fd9-5e67e4787667', 'asdfkj', 'djjdd', 'admin', '2018-07-20 06:08:32');
INSERT INTO `kube_service_version` VALUES ('aee993e2-66e6-4aa0-b3cd-0518f0992a12', 'v1', '', 'admin', '2018-08-10 09:49:54');
INSERT INTO `kube_service_version` VALUES ('b752a59c-ca71-4c64-a264-3e5bbc21e985', 'test', 'test', 'admin', '2018-07-20 06:08:00');
INSERT INTO `kube_service_version` VALUES ('c312762f-c880-44b1-927a-9cc8aeb07e8e', 'v3', 'test', 'admin', '2018-08-13 07:01:23');
INSERT INTO `kube_service_version` VALUES ('c3ab145b-a313-4088-b8e5-28fa98d03842', 'v2', 'ttt', 'admin', '2018-08-13 01:59:18');

-- ----------------------------
-- Table structure for `rms_backend_user`
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user`;
CREATE TABLE `rms_backend_user` (
  `id` varchar(40) NOT NULL,
  `real_name` varchar(255) NOT NULL DEFAULT '',
  `user_name` varchar(255) NOT NULL DEFAULT '',
  `user_pwd` varchar(255) NOT NULL DEFAULT '',
  `is_super` tinyint(1) NOT NULL DEFAULT '0',
  `status` int(11) NOT NULL DEFAULT '0',
  `mobile` varchar(16) NOT NULL DEFAULT '',
  `email` varchar(256) NOT NULL DEFAULT '',
  `avatar` varchar(256) NOT NULL DEFAULT '',
  `user_type` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_backend_user
-- ----------------------------
INSERT INTO `rms_backend_user` VALUES ('1', 'zxbk', 'admin', '21232f297a57a5a743894a0e4a801fc3', '1', '1', '', '', '/static/upload/1.jpg', '0');
INSERT INTO `rms_backend_user` VALUES ('1fb667da-6ed9-4b8e-8899-a2b809337f3f', 'fangjinwei11', 'fangfjinwei_test', '698d51a19d8a121ce581499d7b701668', '1', '1', '15314245011', 'asdfas@163.com', '', '0');
INSERT INTO `rms_backend_user` VALUES ('2', 'fangjinwei', 'fangjinwei', '21232f297a57a5a743894a0e4a801fc3', '0', '1', '15314245011', 'asdfas@163.com', '', '1');
INSERT INTO `rms_backend_user` VALUES ('3', '张三', 'zhangsan', '21232f297a57a5a743894a0e4a801fc3', '1', '1', '', '', '', '0');
INSERT INTO `rms_backend_user` VALUES ('5', '李四asdf', 'lisi', '84d9cfc2f395ce883a41d7ffc1bbcf4e', '0', '0', '15314245011', '1111@163.com', '', '0');
INSERT INTO `rms_backend_user` VALUES ('5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '12121212', 'jikytest', '96e79218965eb72c92a549dd5a330112', '0', '1', '15395105573', 'jikun.zhang@edaibu.net', '', '0');
INSERT INTO `rms_backend_user` VALUES ('6', '张吉坤', 'zjk', '96e79218965eb72c92a549dd5a330112', '1', '1', '15395105573', 'jikun.zhang@edaibu.net', '', '0');
INSERT INTO `rms_backend_user` VALUES ('69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '江 ', 'jxm', '7fa8282ad93047a4d6fe6111c93b308a', '0', '1', '15314245011', 'asdfas@163.com', '', '1');
INSERT INTO `rms_backend_user` VALUES ('766b78a2-af66-4adf-b40e-435c09c5eed4', 'wwww', 'wwww', 'e10adc3949ba59abbe56e057f20f883e', '0', '1', '13933224131', '33123@qq.com', '', '0');
INSERT INTO `rms_backend_user` VALUES ('eabff854-a7fe-4df3-ac78-c1638c9159c2', '123', 'jk-k8s', '96e79218965eb72c92a549dd5a330112', '0', '1', '15395105573', 'jikun.zhang@edaibu.net', '', '0');

-- ----------------------------
-- Table structure for `rms_backend_user_rms_roles`
-- ----------------------------
DROP TABLE IF EXISTS `rms_backend_user_rms_roles`;
CREATE TABLE `rms_backend_user_rms_roles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `rms_backend_user_id` varchar(40) NOT NULL,
  `rms_role_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_backend_user_rms_roles
-- ----------------------------

-- ----------------------------
-- Table structure for `rms_resource`
-- ----------------------------
DROP TABLE IF EXISTS `rms_resource`;
CREATE TABLE `rms_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rtype` int(11) NOT NULL DEFAULT '0',
  `name` varchar(64) NOT NULL DEFAULT '',
  `parent_id` int(11) DEFAULT NULL,
  `seq` int(11) NOT NULL DEFAULT '0',
  `icon` varchar(32) NOT NULL DEFAULT '',
  `url_for` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=103 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_resource
-- ----------------------------
INSERT INTO `rms_resource` VALUES ('7', '1', '权限管理', '8', '100', 'fa fa-balance-scale', '');
INSERT INTO `rms_resource` VALUES ('8', '0', '系统', null, '200', 'fa fa-cog', '');
INSERT INTO `rms_resource` VALUES ('9', '1', '资源管理', '7', '100', 'fa fa-unlock-alt', 'ResourceController.Index');
INSERT INTO `rms_resource` VALUES ('12', '1', '角色管理', '7', '100', 'fa fa-user-secret', 'RoleController.Index');
INSERT INTO `rms_resource` VALUES ('13', '1', '用户管理', '7', '100', 'fa fa-users', 'BackendUserController.Index');
INSERT INTO `rms_resource` VALUES ('21', '0', 'Kubernetes管理', null, '170', 'fa fa-anchor', '');
INSERT INTO `rms_resource` VALUES ('25', '2', '编辑', '9', '100', 'fa fa-pencil', 'ResourceController.Edit');
INSERT INTO `rms_resource` VALUES ('26', '2', '编辑', '13', '100', 'fa fa-pencil', 'BackendUserController.Edit');
INSERT INTO `rms_resource` VALUES ('27', '2', '删除', '9', '100', 'fa fa-trash', 'ResourceController.Delete');
INSERT INTO `rms_resource` VALUES ('29', '2', '删除', '13', '100', 'fa fa-trash', 'BackendUserController.Delete');
INSERT INTO `rms_resource` VALUES ('30', '2', '编辑', '12', '100', 'fa fa-pencil', 'RoleController.Edit');
INSERT INTO `rms_resource` VALUES ('31', '2', '删除', '12', '100', 'fa fa-trash', 'RoleController.Delete');
INSERT INTO `rms_resource` VALUES ('32', '2', '分配资源', '12', '100', 'fa fa-th', 'RoleController.Allocate');
INSERT INTO `rms_resource` VALUES ('35', '1', ' 首页', null, '100', 'fa fa-dashboard', 'HomeController.Index');
INSERT INTO `rms_resource` VALUES ('36', '1', '基础管理', '94', '100', 'fa fa-calculator', '');
INSERT INTO `rms_resource` VALUES ('37', '1', '部署管理', '94', '101', 'fa fa-creative-commons', '');
INSERT INTO `rms_resource` VALUES ('38', '1', '命名空间概况', '94', '99', 'fa fa-bar-chart', 'KubeDashBoardController.Namespaces');
INSERT INTO `rms_resource` VALUES ('39', '0', 'Kubernetes日志管理', null, '180', 'fa fa-rss-square', '');
INSERT INTO `rms_resource` VALUES ('40', '0', '云应用商城', null, '190', 'fa fa-shopping-cart', '');
INSERT INTO `rms_resource` VALUES ('45', '1', '节点', '36', '100', 'fa fa-desktop', 'KubeDashBoardController.Nodes');
INSERT INTO `rms_resource` VALUES ('49', '1', '持久化存储卷', '36', '101', 'fa fa-database', 'KubeDashBoardController.Pvs');
INSERT INTO `rms_resource` VALUES ('50', '1', 'Deployment', '37', '100', 'fa fa-paper-plane-o', 'KubeDashBoardController.Deployment');
INSERT INTO `rms_resource` VALUES ('51', '1', 'Statefulset', '37', '101', 'fa fa-rocket', 'KubeDashBoardController.Statefulset');
INSERT INTO `rms_resource` VALUES ('52', '1', 'Daemonset', '37', '102', 'fa fa-fighter-jet', 'KubeDashBoardController.DaemonSet');
INSERT INTO `rms_resource` VALUES ('56', '1', 'Pod', '37', '103', 'fa fa-star', 'KubeDashBoardController.Pod');
INSERT INTO `rms_resource` VALUES ('69', '1', '日志管理', '39', '100', 'fa fa-search', 'MessagesController.Logs');
INSERT INTO `rms_resource` VALUES ('70', '1', '应用中心', '81', '100', 'fa  fa-ship', 'AppshopController.Apps');
INSERT INTO `rms_resource` VALUES ('71', '1', '添加应用', '81', '101', 'fa fa-plus-square', '');
INSERT INTO `rms_resource` VALUES ('72', '1', '部署管理', '81', '102', 'fa fa-cutlery', '');
INSERT INTO `rms_resource` VALUES ('74', '1', '部署服务', '93', '90', 'fa fa-birthday-cake', 'KubeController.Deploy');
INSERT INTO `rms_resource` VALUES ('75', '0', '基础环境管理', null, '110', '', '');
INSERT INTO `rms_resource` VALUES ('76', '1', 'HOST管理', '80', '100', 'fa fa-tv', 'BaseclusterController.Hosts');
INSERT INTO `rms_resource` VALUES ('78', '1', '集群管理', '80', '101', 'fa fa-recycle', 'BaseclusterController.Cluster');
INSERT INTO `rms_resource` VALUES ('79', '1', '环境设置', '80', '102', 'fa fa-share-alt-square', 'BaseclusterController.ClusterSetup');
INSERT INTO `rms_resource` VALUES ('80', '1', '基础环境管理', '75', '100', 'fa fa-bank', '');
INSERT INTO `rms_resource` VALUES ('81', '1', '云商城', '40', '100', 'fa fa-shopping-cart', '');
INSERT INTO `rms_resource` VALUES ('82', '0', '基础信息管理', null, '120', '', '');
INSERT INTO `rms_resource` VALUES ('83', '1', '命名空间管理', '84', '100', 'fa fa-delicious', 'BaseclusterSetupController.NameSpaceSetup');
INSERT INTO `rms_resource` VALUES ('84', '1', '基础信息管理', '82', '100', 'fa fa-road', '');
INSERT INTO `rms_resource` VALUES ('86', '1', '集群角色-集群', '84', '102', 'fa fa-user-plus', 'BaseclusterSetupController.ClusterRole');
INSERT INTO `rms_resource` VALUES ('87', '1', '用户组授权-集群', '84', '103', 'fa fa-group', 'BaseclusterSetupController.GrantUserGroup');
INSERT INTO `rms_resource` VALUES ('89', '1', 'SA', '84', '105', 'fa fa-mortar-board', 'BaseclusterSetupController.Sa');
INSERT INTO `rms_resource` VALUES ('90', '0', '应用服务管理', null, '160', '', '');
INSERT INTO `rms_resource` VALUES ('91', '1', 'Proxy管理', '93', '100', 'fa fa-rss-square', 'KubeController.DeployProxy');
INSERT INTO `rms_resource` VALUES ('92', '1', 'Ingress管理', '93', '98', 'fa fa-shopping-cart', 'KubeController.DeployIngress');
INSERT INTO `rms_resource` VALUES ('93', '1', '应用服务管理', '90', '100', 'fa fa-rss-square', '');
INSERT INTO `rms_resource` VALUES ('94', '1', 'Kubernetes管理', '21', '100', 'fa fa-anchor', '');
INSERT INTO `rms_resource` VALUES ('95', '1', 'Images镜像管理', '93', '101', 'fa fa-cubes', 'KubeController.Images');
INSERT INTO `rms_resource` VALUES ('96', '1', '存储类', '36', '102', 'fa fa-download', 'KubeDashBoardController.StorageClass');
INSERT INTO `rms_resource` VALUES ('99', '1', 'Service', '101', '100', 'fa fa-object-group', 'KubeDashBoardController.Services');
INSERT INTO `rms_resource` VALUES ('100', '1', 'CronJob', '37', '103', 'fa fa-calendar-check-o', 'KubeDashBoardController.CronJob');
INSERT INTO `rms_resource` VALUES ('101', '1', '服务发现与负载均衡', '94', '103', 'fa fa-sitemap', '');
INSERT INTO `rms_resource` VALUES ('102', '1', 'Ingress', '101', '101', 'fa  fa-share-alt', '');

-- ----------------------------
-- Table structure for `rms_role`
-- ----------------------------
DROP TABLE IF EXISTS `rms_role`;
CREATE TABLE `rms_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL DEFAULT '',
  `seq` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_role
-- ----------------------------
INSERT INTO `rms_role` VALUES ('22', '超级管理员', '20');
INSERT INTO `rms_role` VALUES ('24', '角色管理员', '10');
INSERT INTO `rms_role` VALUES ('25', '普通管理员', '5');
INSERT INTO `rms_role` VALUES ('26', 'test', '33');

-- ----------------------------
-- Table structure for `rms_role_backenduser_rel`
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_backenduser_rel`;
CREATE TABLE `rms_role_backenduser_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `backend_user_id` varchar(40) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_role_backenduser_rel
-- ----------------------------
INSERT INTO `rms_role_backenduser_rel` VALUES ('68', '22', '1', '2018-06-15 06:02:04');
INSERT INTO `rms_role_backenduser_rel` VALUES ('84', '25', '3', '2018-07-03 08:00:57');
INSERT INTO `rms_role_backenduser_rel` VALUES ('131', '25', '1fb667da-6ed9-4b8e-8899-a2b809337f3f', '2018-07-11 02:22:14');
INSERT INTO `rms_role_backenduser_rel` VALUES ('132', '24', '1fb667da-6ed9-4b8e-8899-a2b809337f3f', '2018-07-11 02:22:14');
INSERT INTO `rms_role_backenduser_rel` VALUES ('137', '24', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '2018-07-11 03:33:30');
INSERT INTO `rms_role_backenduser_rel` VALUES ('138', '24', '6', '2018-07-11 05:30:27');
INSERT INTO `rms_role_backenduser_rel` VALUES ('139', '22', '6', '2018-07-11 05:30:27');
INSERT INTO `rms_role_backenduser_rel` VALUES ('141', '25', '5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '2018-07-11 05:45:40');
INSERT INTO `rms_role_backenduser_rel` VALUES ('142', '24', '5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '2018-07-11 05:45:40');
INSERT INTO `rms_role_backenduser_rel` VALUES ('143', '22', '5e3605af-1b0f-45fb-97e1-a158ce7cfd5b', '2018-07-11 05:45:40');
INSERT INTO `rms_role_backenduser_rel` VALUES ('144', '25', '5', '2018-08-16 02:12:06');
INSERT INTO `rms_role_backenduser_rel` VALUES ('146', '25', '766b78a2-af66-4adf-b40e-435c09c5eed4', '2018-08-16 03:25:46');
INSERT INTO `rms_role_backenduser_rel` VALUES ('147', '24', '766b78a2-af66-4adf-b40e-435c09c5eed4', '2018-08-16 03:25:46');
INSERT INTO `rms_role_backenduser_rel` VALUES ('148', '22', '766b78a2-af66-4adf-b40e-435c09c5eed4', '2018-08-16 03:25:46');
INSERT INTO `rms_role_backenduser_rel` VALUES ('149', '26', '766b78a2-af66-4adf-b40e-435c09c5eed4', '2018-08-16 03:25:46');
INSERT INTO `rms_role_backenduser_rel` VALUES ('152', '26', 'eabff854-a7fe-4df3-ac78-c1638c9159c2', '2018-08-17 07:13:43');

-- ----------------------------
-- Table structure for `rms_role_resource_rel`
-- ----------------------------
DROP TABLE IF EXISTS `rms_role_resource_rel`;
CREATE TABLE `rms_role_resource_rel` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL,
  `resource_id` int(11) NOT NULL,
  `created` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=565 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rms_role_resource_rel
-- ----------------------------
INSERT INTO `rms_role_resource_rel` VALUES ('429', '25', '21', '2017-12-19 06:40:04');
INSERT INTO `rms_role_resource_rel` VALUES ('451', '22', '35', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('452', '22', '21', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('453', '22', '36', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('454', '22', '45', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('458', '22', '49', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('459', '22', '37', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('460', '22', '50', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('461', '22', '51', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('462', '22', '52', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('466', '22', '38', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('468', '22', '56', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('484', '22', '39', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('485', '22', '69', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('486', '22', '40', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('487', '22', '70', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('488', '22', '71', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('489', '22', '72', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('490', '22', '8', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('491', '22', '7', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('492', '22', '9', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('493', '22', '25', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('494', '22', '27', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('495', '22', '12', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('496', '22', '30', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('497', '22', '31', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('498', '22', '32', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('499', '22', '13', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('500', '22', '26', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('501', '22', '29', '2018-06-15 10:27:04');
INSERT INTO `rms_role_resource_rel` VALUES ('502', '24', '35', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('503', '24', '75', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('504', '24', '80', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('505', '24', '76', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('506', '24', '78', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('507', '24', '79', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('508', '24', '82', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('509', '24', '84', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('510', '24', '83', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('511', '24', '86', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('512', '24', '87', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('513', '24', '89', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('514', '24', '21', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('515', '24', '74', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('516', '24', '36', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('517', '24', '45', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('521', '24', '49', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('522', '24', '37', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('523', '24', '50', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('524', '24', '51', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('525', '24', '52', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('529', '24', '38', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('531', '24', '56', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('546', '24', '39', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('547', '24', '69', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('548', '24', '40', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('549', '24', '81', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('550', '24', '70', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('551', '24', '71', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('552', '24', '72', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('553', '24', '8', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('554', '24', '7', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('555', '24', '9', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('556', '24', '25', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('557', '24', '27', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('558', '24', '12', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('559', '24', '30', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('560', '24', '31', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('561', '24', '32', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('562', '24', '13', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('563', '24', '26', '2018-07-11 02:31:49');
INSERT INTO `rms_role_resource_rel` VALUES ('564', '24', '29', '2018-07-11 02:31:49');

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(40) NOT NULL,
  `user_name` varchar(20) NOT NULL,
  `pass_word` varchar(20) NOT NULL,
  `type` int(11) NOT NULL COMMENT '管理员/普通用户',
  `create_user` varchar(40) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `remark` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户';

-- ----------------------------
-- Records of user
-- ----------------------------

-- ----------------------------
-- Table structure for `user_and_group`
-- ----------------------------
DROP TABLE IF EXISTS `user_and_group`;
CREATE TABLE `user_and_group` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `group_id` varchar(40) NOT NULL COMMENT '用户组id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和用户组中间表';

-- ----------------------------
-- Records of user_and_group
-- ----------------------------
INSERT INTO `user_and_group` VALUES ('88fb42d8-46e3-40b1-8ca2-cbcf507bdc3c', '6', '12bf01b6-cfb5-442d-b063-993c74849bfe');

-- ----------------------------
-- Table structure for `user_group`
-- ----------------------------
DROP TABLE IF EXISTS `user_group`;
CREATE TABLE `user_group` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `group_name` varchar(40) NOT NULL COMMENT '用户组名称',
  `cluster_id` varchar(40) NOT NULL COMMENT '集群id',
  `create_user` varchar(40) DEFAULT NULL COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `remark` varchar(200) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='k8s用户组';

-- ----------------------------
-- Records of user_group
-- ----------------------------
INSERT INTO `user_group` VALUES ('12bf01b6-cfb5-442d-b063-993c74849bfe', '123', '1', '1', '2018-07-09 06:28:04', '2222233');
INSERT INTO `user_group` VALUES ('3a923331-cb1d-4b0d-bbdd-d595d404e59e', 'fkw_test111', '1', '1', '2018-07-11 02:23:40', 'ddddd111');
INSERT INTO `user_group` VALUES ('3ea7833c-78be-46c6-b0ab-e5f345cc89a1', 'dltest11', '1', '1', '2018-08-16 10:20:31', '');
INSERT INTO `user_group` VALUES ('4d960657-378f-46de-b156-12c7b34edffd', '3333', '1', '1', '2018-07-10 07:04:59', '是的发送到');
INSERT INTO `user_group` VALUES ('859302da-7cf4-4143-9a8a-17a2575c1444', 'zjk', '96724a68-d10a-49f6-b852-b3c1d053e238', '1', '2018-08-17 06:27:18', 'zjk');
INSERT INTO `user_group` VALUES ('fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849', '2222', '1', '1', '2018-07-11 06:28:44', '是的发的');

-- ----------------------------
-- Table structure for `user2group`
-- ----------------------------
DROP TABLE IF EXISTS `user2group`;
CREATE TABLE `user2group` (
  `id` varchar(40) NOT NULL COMMENT '主键',
  `user_id` varchar(40) NOT NULL COMMENT '用户id',
  `group_id` varchar(40) NOT NULL COMMENT '用户组id',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户和用户组中间表';

-- ----------------------------
-- Records of user2group
-- ----------------------------
INSERT INTO `user2group` VALUES ('09515177-07bc-4ca1-ba4a-d9f707356470', 'b7ea94d6-4b9b-4a1d-ae57-1d15864203c6', '859302da-7cf4-4143-9a8a-17a2575c1444');
INSERT INTO `user2group` VALUES ('4c3f5107-cec8-47e8-804c-bedfb3d7502e', '1fb667da-6ed9-4b8e-8899-a2b809337f3f', 'fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849');
INSERT INTO `user2group` VALUES ('5d06ac9c-aa04-43b5-b78b-52f610615b65', '69d4a6c6-ebf0-46fb-894a-635d0e23f4a3', '12bf01b6-cfb5-442d-b063-993c74849bfe');
INSERT INTO `user2group` VALUES ('6ba7052b-ac89-4769-91e1-fd930e0461c8', '4214e1ec-d163-4d28-b9d8-5cc8c3248b8c', '4d960657-378f-46de-b156-12c7b34edffd');
INSERT INTO `user2group` VALUES ('789eb5ae-417c-4a6b-a604-a0c35dbf9807', '4214e1ec-d163-4d28-b9d8-5cc8c3248b8c', 'fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849');
INSERT INTO `user2group` VALUES ('79a0bf44-150d-48ae-9285-8c6d11d280fa', 'eabff854-a7fe-4df3-ac78-c1638c9159c2', '859302da-7cf4-4143-9a8a-17a2575c1444');
INSERT INTO `user2group` VALUES ('80963e8f-5936-4317-b0e3-492a6a1c89e4', '6', '4d960657-378f-46de-b156-12c7b34edffd');
INSERT INTO `user2group` VALUES ('ae1051e1-0da8-43ef-ae90-a4d94c87d98d', '766b78a2-af66-4adf-b40e-435c09c5eed4', '3a923331-cb1d-4b0d-bbdd-d595d404e59e');
INSERT INTO `user2group` VALUES ('af9fac4e-3593-4f8b-8670-f89aed548381', '6', 'fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849');
INSERT INTO `user2group` VALUES ('c93378dd-b079-4a11-9a3f-bfec62c79bda', '1fb667da-6ed9-4b8e-8899-a2b809337f3f', '12bf01b6-cfb5-442d-b063-993c74849bfe');
INSERT INTO `user2group` VALUES ('f1b47794-f169-485a-ada0-3b3ad4a846da', '766b78a2-af66-4adf-b40e-435c09c5eed4', 'fd1f4228-5f2f-4c0b-b0e2-5bca4f3fd849');
INSERT INTO `user2group` VALUES ('f34bc699-8028-4ae6-8dec-caaec16d9147', '4214e1ec-d163-4d28-b9d8-5cc8c3248b8c', '12bf01b6-cfb5-442d-b063-993c74849bfe');
INSERT INTO `user2group` VALUES ('fc8b4720-3501-49a9-bf62-f2161770c780', '766b78a2-af66-4adf-b40e-435c09c5eed4', '12bf01b6-cfb5-442d-b063-993c74849bfe');
INSERT INTO `user2group` VALUES ('fe018b92-bc83-4903-b313-f353e03288c2', '766b78a2-af66-4adf-b40e-435c09c5eed4', '4d960657-378f-46de-b156-12c7b34edffd');

-- ----------------------------
-- Procedure structure for `proc_EnvUserCluster`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_EnvUserCluster`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_EnvUserCluster`(IN  _userId VARCHAR(40))
BEGIN

  DECLARE uIsSpuer int;

  select is_super into uIsSpuer from rms_backend_user where id =_userId;
  
  -- select uIsSpuer;

	IF (uIsSpuer =1) THEN
			select c.id,c.`name`,c.remark,IFNULL(euc.id,0) selected from kube_cluster c left JOIN (select * from kube_env_user_cluster where user_id =_userId) euc on (c.id = euc.cluster_id);
	ELSE
			select c.id,c.`name`,c.remark,IFNULL(euc.id,0) selected from kube_auth_user_cluster auc left join kube_cluster c on(auc.cluster_id = c.id) left JOIN (select * from kube_env_user_cluster where user_id =_userId) euc on(auc.cluster_id = euc.cluster_id) where auc.user_id =_userId and auc.user_type=0 ;
  END IF;

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_EnvUserNameSpace`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_EnvUserNameSpace`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_EnvUserNameSpace`(IN  _userId VARCHAR(40),IN  _clusterId VARCHAR(40))
BEGIN

  DECLARE uIsSpuer int;
  DECLARE uType int;


  select is_super into uIsSpuer  from rms_backend_user where id =_userId;
  select user_type into uType  from rms_backend_user where id =_userId;
  
  -- select uIsSpuer;

	IF (uIsSpuer =1) THEN
			select ns.id,ns.`name`,ns.remark,IFNULL(euns.namespace_id,0) selected from kube_namespace ns left join (select * from  kube_env_user_namespace where user_id =_userId) euns on (ns.id = euns.namespace_id) where ns.cluster_id = _clusterId;
	ELSEIF (uType =1) THEN
			select ns.id,ns.`name`,ns.remark,IFNULL(euns.namespace_id,0) selected from kube_namespace ns left join (select * from  kube_env_user_namespace where user_id =_userId) euns on (ns.id = euns.namespace_id) where ns.cluster_id = _clusterId;
	ELSE
			select ns.id,ns.`name`,ns.remark,IFNULL(euns.namespace_id,0) selected from kube_auth_user_namespace auns left join kube_namespace ns on (auns.namespace_id = ns.id) left join (select * from  kube_env_user_namespace where user_id =_userId) euns on (ns.id = euns.namespace_id) where auns.user_type =0 and auns.user_id =_userId and ns.cluster_id = _clusterId;
  END IF;

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubeAuthUserNameSpaceQPL`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubeAuthUserNameSpaceQPL`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubeAuthUserNameSpaceQPL`(IN  _iId VARCHAR(40) ,
IN  _userId VARCHAR(40) ,
IN  _iUserName VARCHAR(40) ,
IN  _clusterId VARCHAR(40) ,
IN  _iClusterName VARCHAR(40) ,
IN  _nameSpaceId VARCHAR(40) ,
IN  _iNameSpaceName VARCHAR(40) ,
IN  _sort VARCHAR(40) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN
 set @sqlWhere = CONCAT_WS(' ','select auns.*,ns.`name` namespace_name,ct.`name` cluster_name,bu.real_name user_name  from kube_auth_user_namespace auns left JOIN kube_namespace ns on(auns.namespace_id = ns.id) left JOIN kube_cluster ct on (auns.cluster_id = ct.id) left JOIN rms_backend_user bu on (auns.user_id = bu.id)',
	'where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and auns.id = \'', userREPLACE(_iId),'\'') END,
  CASE _userId WHEN '' THEN '' ELSE CONCAT(' and auns.user_id = \'', userREPLACE(_userId),'\'') END,
  CASE _clusterId WHEN '' THEN '' ELSE CONCAT(' and auns.cluster_id = \'', userREPLACE(_clusterId) ,'\'') END,
  CASE _nameSpaceId WHEN '' THEN '' ELSE CONCAT(' and auns.namespace_id = \'', userREPLACE(_nameSpaceId) ,'\'') END);

 	set @sqlText = CONCAT_WS(' ',@sqlWhere,
 	' order by ',
 	 CASE
       IFNULL(_sort, '') 
       WHEN '' 
       THEN 'id ' 
       ELSE CONCAT(_sort,' ')
     END,
 	 CASE
       IFNULL(_order, '') 
       WHEN '' 
       THEN ' desc ' 
       ELSE CONCAT(_order,' ') 
     END,
 	  ' limit ',_offset,' , ',CASE IFNULL(_limit,0) WHEN 0 THEN 2147483647 ELSE _limit END
 	);


  
  PREPARE strsql FROM @sqlText ;
  EXECUTE strsql ;

 

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubeAuthUserNameSpaceQPL_Test`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubeAuthUserNameSpaceQPL_Test`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubeAuthUserNameSpaceQPL_Test`(IN  _iId VARCHAR(40) ,
IN  _userId VARCHAR(40) ,
IN  _iUserName VARCHAR(40) ,
IN  _clusterId VARCHAR(40) ,
IN  _iClusterName VARCHAR(40) ,
IN  _nameSpaceId VARCHAR(40) ,
IN  _iNameSpaceName VARCHAR(40) ,
IN  _sort VARCHAR(40) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN
 set @sqlWhere = CONCAT_WS(' ','select auns.*,ns.`name` namespace_name,ct.`name` cluster_name,bu.real_name user_name  from kube_auth_user_namespace auns left JOIN kube_namespace ns on(auns.namespace_id = ns.id) left JOIN kube_cluster ct on (auns.cluster_id = ct.id) left JOIN rms_backend_user bu on (auns.user_id = bu.id)',
	'where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and auns.id = \'', userREPLACE(_iId),'\'') END,
  CASE _userId WHEN '' THEN '' ELSE CONCAT(' and auns.user_id = \'',  userREPLACE(_userId),'\'') END,
  CASE _clusterId WHEN '' THEN '' ELSE CONCAT(' and auns.cluster_id = \'', userREPLACE(_clusterId) ,'\'') END,
  CASE _nameSpaceId WHEN '' THEN '' ELSE CONCAT(' and auns.namespace_id = \'', userREPLACE(_nameSpaceId) ,'\'') END);

 	set @sqlText = CONCAT_WS(' ',@sqlWhere,
 	' order by ',
 	 CASE
       IFNULL(_sort, '') 
       WHEN '' 
       THEN 'id ' 
       ELSE CONCAT(_sort,' ')
     END,
 	 CASE
       IFNULL(_order, '') 
       WHEN '' 
       THEN ' desc ' 
       ELSE CONCAT(_order,' ') 
     END,
 	  ' limit ',_offset,' , ',CASE IFNULL(_limit,0) WHEN 0 THEN 2147483647 ELSE _limit END
 	);


  
  PREPARE strsql FROM @sqlText ;
  EXECUTE strsql ;

 

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubeAuthUserNameSpaceQTC`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubeAuthUserNameSpaceQTC`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubeAuthUserNameSpaceQTC`(IN  _iId VARCHAR(40) ,
IN  _userId VARCHAR(40) ,
IN  _iUserName VARCHAR(40) ,
IN  _clusterId VARCHAR(40) ,
IN  _iClusterName VARCHAR(40) ,
IN  _nameSpaceId VARCHAR(40) ,
IN  _iNameSpaceName VARCHAR(40) ,
IN  _sort VARCHAR(40) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN
 set @sqlWhere = CONCAT_WS(' ','select auns.*,ns.`name` namespace_name,ct.`name` clusert_name,bu.real_name user_name  from kube_auth_user_namespace auns left JOIN kube_namespace ns on(auns.namespace_id = ns.id) left JOIN kube_cluster ct on (auns.cluster_id = ct.id) left JOIN rms_backend_user bu on (auns.user_id = bu.id)',
	'where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and auns.id = \'', userREPLACE(_iId),'\'') END,
  CASE _userId WHEN '' THEN '' ELSE CONCAT(' and auns.user_id = \'', userREPLACE(_userId),'\'') END,
  CASE _clusterId WHEN '' THEN '' ELSE CONCAT(' and auns.cluster_id = \'', userREPLACE(_clusterId) ,'\'') END,
  CASE _nameSpaceId WHEN '' THEN '' ELSE CONCAT(' and auns.namespace_id = \'', userREPLACE(_nameSpaceId) ,'\'') END);

 	set @sqlCount =CONCAT('select COUNT(*) search_count from (',@sqlWhere,') getCount');
 
	PREPARE strsqlCount FROM @sqlCount ;
  EXECUTE strsqlCount ;

 

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubeNameSpaceQPL`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubeNameSpaceQPL`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubeNameSpaceQPL`(IN  _iId VARCHAR(40) ,
IN  _iName VARCHAR(40), IN  _iClusterId VARCHAR(40),
IN  _sort VARCHAR(160) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN

  set @sqlWhere = CONCAT_WS(' ','select ns.*,cr.`name` as cluster_name from kube_namespace ns left JOIN kube_cluster cr on (ns.cluster_id = cr.id)',
	' where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and ns.id = \'', userREPLACE(_iId),'\'') END,
	CASE _iClusterId  WHEN '' THEN '' ELSE CONCAT(' and ns.cluster_id = \'', userREPLACE(_iClusterId) ,'\'') END,
	CASE _iName  WHEN '' THEN '' ELSE CONCAT(' and ns.name like \'%', userREPLACE(_iName) ,'%\'') END);

	set @sqlText = CONCAT_WS(' ',@sqlWhere,
	' order by ',
	 CASE
      IFNULL(_sort, '') 
      WHEN '' 
      THEN 'id ' 
      ELSE CONCAT(_sort,' ')
    END,
	 CASE
      IFNULL(_order, '') 
      WHEN '' 
      THEN ' desc ' 
      ELSE CONCAT(_order,' ') 
    END,
	  ' limit ',_offset,' , ',CASE IFNULL(_limit,0) WHEN 0 THEN 2147483647 ELSE _limit END
	);

	set @sqlCount =CONCAT('select COUNT(*) count from (',@sqlWhere,') getCount');
  
  PREPARE strsql FROM @sqlText ;
  EXECUTE strsql ;

	PREPARE strsqlCount FROM @sqlCount ;
  EXECUTE strsqlCount ;

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubeNameSpaceQTC`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubeNameSpaceQTC`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubeNameSpaceQTC`(IN  _iId VARCHAR(40) ,
IN  _iName VARCHAR(40), IN  _iClusterId VARCHAR(40) ,
IN  _sort VARCHAR(160) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN
  set @sqlWhere = CONCAT_WS(' ','select ns.*,cr.`name` as cluster_name from kube_namespace ns left JOIN kube_cluster cr on (ns.cluster_id = cr.id)',
	' where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and ns.id = \'', userREPLACE(_iId),'\'') END,
  CASE _iClusterId  WHEN '' THEN '' ELSE CONCAT(' and ns.cluster_id = \'', userREPLACE(_iClusterId) ,'\'') END,
	CASE _iName  WHEN '' THEN '' ELSE CONCAT(' and ns.name like \'%', userREPLACE(_iName) ,'%\'') END);

	set @sqlCount =CONCAT('select COUNT(*) search_count from (',@sqlWhere,') getCount');
 
	PREPARE strsqlCount FROM @sqlCount ;
  EXECUTE strsqlCount ;
--
END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubePublishServiceQPL`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubePublishServiceQPL`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubePublishServiceQPL`(IN  _iId VARCHAR(40) ,IN  _iStype VARCHAR(40) ,
IN  _iName VARCHAR(100), IN  _iServiceId VARCHAR(40) ,IN  _iServiceName VARCHAR(40) ,IN  _iClusterId VARCHAR(40),
IN  _iNamespaceId VARCHAR(40),IN  _iNamespaceName VARCHAR(40),IN  _iRemark VARCHAR(40),IN  _sort VARCHAR(160) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN

  -- set @sqlWhere = CONCAT_WS(' ','select pp.*,s.`name` server_name,sp.container_port,sp.service_port, c.`name` cluster_name ,ns.`name` namespace_name  from kube_publish_proxy pp LEFT JOIN kube_service s on (pp.service_id = s.id) left join kube_service_port sp on(pp.port_id = sp.id) left JOIN kube_cluster c on(pp.cluster_id = c.id) left join kube_namespace ns on(pp.namespace_id = ns.id)',
  set @sqlWhere = CONCAT_WS(' ','select ps.*,s.`name` server_name, c.`name` cluster_name ,ns.`name` namespace_name  from kube_publish_service ps LEFT JOIN kube_service s on (ps.service_id = s.id) left JOIN kube_cluster c on(ps.cluster_id = c.id) left join kube_namespace ns on(ps.namespace_id = ns.id)',
	' where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and ps.id = \'', userREPLACE(_iId),'\'') END,
	CASE _iStype WHEN '' THEN '' ELSE CONCAT(' and ps.stype = \'', userREPLACE(_iStype ),'\'') END,
	CASE _iName    WHEN '' THEN '' ELSE CONCAT(' and ps.name like \'%', userREPLACE(_iName) ,'%\'') END,
	CASE _iServiceId   WHEN '' THEN '' ELSE CONCAT(' and ps.service_id = \'', userREPLACE(_iServiceId) ,'\'') END,
	CASE _iServiceName   WHEN '' THEN '' ELSE CONCAT(' and s.name like \'%', userREPLACE(_iServiceName) ,'%\'') END,
	CASE _iClusterId  WHEN '' THEN '' ELSE CONCAT(' and ps.cluster_id = \'', userREPLACE(_iClusterId) ,'\'') END,
	CASE _iNamespaceId    WHEN '' THEN '' ELSE CONCAT(' and ps.namespace_id = \'', userREPLACE(_iNamespaceId ) ,'\'') END,
	CASE _iNamespaceName    WHEN '' THEN '' ELSE CONCAT(' and ns.name like \'%', userREPLACE(_iNamespaceName ) ,'%\'') END,
	CASE _iRemark   WHEN '' THEN '' ELSE CONCAT(' and ps.remark like \'%', userREPLACE(_iRemark ) ,'%\'') END);

	set @sqlText = CONCAT_WS(' ',@sqlWhere,
	' order by ',
	 CASE
      IFNULL(_sort, '') 
      WHEN '' 
      THEN 'id ' 
      ELSE CONCAT(_sort,' ')
    END,
	 CASE
      IFNULL(_order, '') 
      WHEN '' 
      THEN ' desc ' 
      ELSE CONCAT(_order,' ') 
    END,
	  ' limit ',_offset,' , ',CASE IFNULL(_limit,0) WHEN 0 THEN 2147483647 ELSE _limit END
	);

	-- select @sqlText;

	set @sqlCount =CONCAT('select COUNT(*) count from (',@sqlWhere,') getCount');
  
  PREPARE strsql FROM @sqlText ;
  EXECUTE strsql ;

	PREPARE strsqlCount FROM @sqlCount ;
  EXECUTE strsqlCount ;

END
;;
DELIMITER ;

-- ----------------------------
-- Procedure structure for `proc_KubePublishServiceQTC`
-- ----------------------------
DROP PROCEDURE IF EXISTS `proc_KubePublishServiceQTC`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` PROCEDURE `proc_KubePublishServiceQTC`(IN  _iId VARCHAR(40),IN  _iStype VARCHAR(40) ,
IN  _iName VARCHAR(100), IN  _iServiceId VARCHAR(40) ,IN  _iServiceName VARCHAR(40) ,IN  _iClusterId VARCHAR(40),
IN  _iNamespaceId VARCHAR(40),IN  _iNamespaceName VARCHAR(40),IN  _iRemark VARCHAR(40),IN  _sort VARCHAR(160) ,
IN  _order VARCHAR(12) ,
IN  _offset INT,
IN  _limit INT)
BEGIN

  -- set @sqlWhere = CONCAT_WS(' ','select pp.*,s.`name` server_name,sp.container_port,sp.service_port, c.`name` cluster_name ,ns.`name` namespace_name  from kube_publish_proxy pp LEFT JOIN kube_service s on (pp.service_id = s.id) left join kube_service_port sp on(pp.port_id = sp.id) left JOIN kube_cluster c on(pp.cluster_id = c.id) left join kube_namespace ns on(pp.namespace_id = ns.id)',
  set @sqlWhere = CONCAT_WS(' ','select ps.*,s.`name` server_name, c.`name` cluster_name ,ns.`name` namespace_name  from kube_publish_service ps LEFT JOIN kube_service s on (ps.service_id = s.id) left JOIN kube_cluster c on(ps.cluster_id = c.id) left join kube_namespace ns on(ps.namespace_id = ns.id)',
	' where 1=1',
	CASE _iId WHEN '' THEN '' ELSE CONCAT(' and ps.id = \'', userREPLACE(_iId),'\'') END,
	CASE _iStype WHEN '' THEN '' ELSE CONCAT(' and ps.stype = \'', userREPLACE(_iStype ),'\'') END,
	CASE _iName WHEN '' THEN '' ELSE CONCAT(' and ps.name like \'%', userREPLACE(_iName) ,'%\'') END,
	CASE _iServiceId   WHEN '' THEN '' ELSE CONCAT(' and ps.service_id = \'', userREPLACE(_iServiceId) ,'\'') END,
	CASE _iServiceName   WHEN '' THEN '' ELSE CONCAT(' and s.name like \'%', userREPLACE(_iServiceName) ,'%\'') END,
	CASE _iClusterId  WHEN '' THEN '' ELSE CONCAT(' and ps.cluster_id = \'', userREPLACE(_iClusterId) ,'\'') END,
	CASE _iNamespaceId    WHEN '' THEN '' ELSE CONCAT(' and ps.namespace_id = \'', userREPLACE(_iNamespaceId ) ,'\'') END,
	CASE _iNamespaceName    WHEN '' THEN '' ELSE CONCAT(' and ns.name like \'%', userREPLACE(_iNamespaceName ) ,'%\'') END,
	CASE _iRemark   WHEN '' THEN '' ELSE CONCAT(' and ps.remark like \'%', userREPLACE(_iRemark ) ,'%\'') END);

	set @sqlCount =CONCAT('select COUNT(*) search_count from (',@sqlWhere,') getCount');

	PREPARE strsqlCount FROM @sqlCount ;
  EXECUTE strsqlCount ;

END
;;
DELIMITER ;

-- ----------------------------
-- Function structure for `userREPLACE`
-- ----------------------------
DROP FUNCTION IF EXISTS `userREPLACE`;
DELIMITER ;;
CREATE DEFINER=`root`@`%` FUNCTION `userREPLACE`(v VARCHAR(400)) RETURNS varchar(400) CHARSET utf8
BEGIN
  -- set REPLACE(v,'''','')
 
  -- 返回函数处理结果
  RETURN REPLACE(REPLACE(REPLACE(v,'--',''),'/',''),'''','');
END
;;
DELIMITER ;
