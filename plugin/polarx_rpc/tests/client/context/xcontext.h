/*
 * Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License, version 2.0,
 * as published by the Free Software Foundation.
 *
 * This program is also distributed with certain software (including
 * but not limited to OpenSSL) that is licensed under separate terms,
 * as designated in a particular file or component or in included license
 * documentation.  The authors of MySQL hereby grant you an additional
 * permission to link the program and your derivative works with the
 * separately licensed software that they have included with MySQL.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License, version 2.0, for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */

// MySQL DB access module, for use by plugins and others
// For the module that implements interactive DB functionality see mod_db

#ifndef PLUGIN_X_CLIENT_CONTEXT_XCONTEXT_H_
#define PLUGIN_X_CLIENT_CONTEXT_XCONTEXT_H_

#include <cstring>
#include <string>
#include <vector>

#include "../mysqlxclient/xerror.h"
#include "../mysqlxclient/xprotocol.h"
#include "xconnection_config.h"

namespace xcl {

enum class Auth {
  k_auto,
  k_auto_fallback,
  k_auto_from_capabilities,
  k_mysql41,
  //  k_plain,
  //  k_sha256_memory
};

class Context {
 public:
  Connection_config m_connection_config;
  bool m_consume_all_notices{true};
  XProtocol::Client_id m_client_id{XCL_CLIENT_ID_NOT_VALID};
  XError m_global_error;
  // Default value equal to lenght of DateTime when no time part is present
  uint32_t m_datetime_length_discriminator = 10;
  Internet_protocol m_internet_protocol{Internet_protocol::Any};
  std::vector<Auth> m_use_auth_methods;
};

}  // namespace xcl

#endif  // PLUGIN_X_CLIENT_CONTEXT_XCONTEXT_H_
